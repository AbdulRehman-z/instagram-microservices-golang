package api

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	db "github.com/AbdulRehman-z/instagram-microservices/auth_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/types"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/util"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/worker"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"github.com/lib/pq"
)

func (server *Server) RegisterUser(c *fiber.Ctx) error {
	var req types.RegisterUserReqParams
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "internal err")
	}

	args := db.RegisterUserTXParams{
		RegisterUserParams: db.RegisterUserParams{
			Email:          req.Email,
			HashedPassword: hashedPassword,
		},
		AfterRegister: func(user db.User) error {
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.CriticaLQueue),
			}

			payload := &worker.PayloadSendVerificationEmail{
				Email: user.Email,
			}
			err := server.taskDistributor.TaskSendSignupEmail(c.UserContext(), payload, opts...)
			if err != nil {
				return fmt.Errorf("cannot send verification email: %w", err)
			}
			return nil
		},
	}

	userTx, err := server.store.RegisterUserTx(c.UserContext(), args)
	if err != nil {
		if pqError, ok := err.(*pq.Error); ok {
			if pqError.Code.Name() == "unique_violation" {
				return fiber.NewError(fiber.StatusConflict, err.Error())
			}
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := types.RegisterUserResParams{
		Email:             userTx.User.Email,
		PasswordChangedAt: userTx.User.PasswordChangedAt,
		PasswordCreatedAt: userTx.User.CreatedAt,
	}
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"data": response,
	})
}

func (server *Server) LoginUser(c *fiber.Ctx) error {
	var req types.LoginUserReqParams
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := server.store.GetUser(context.Background(), req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if !user.IsEmailVerified {
		return fiber.NewError(fiber.StatusBadRequest, "email not verified")
	}

	if err := util.ComparePassword(req.Password, user.HashedPassword); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid credentials")
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(user.Email, server.Config.ACCESS_TOKEN_DURATION)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "internal err")
	}

	fmt.Println(accessToken, accessPayload)

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(user.Email, server.Config.REFRESH_TOKEN_DURATION)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "internal err")
	}

	res := server.rStore.HSet(context.Background(), refreshPayload.Id.String(), "email", refreshPayload.Email,
		"refreshToken", refreshToken, "expiresAt", refreshPayload.ExpiredAt)
	fmt.Println(res)
	return nil
}

func (server *Server) ChangePassword(c *fiber.Ctx) error {
	var req types.ChangePasswordReqParams
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	hashedPassword, err := util.HashPassword(req.NewPassword)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "internal err")
	}

	userTx, err := server.store.ChangePasswordTx(c.Context(), db.ChangePasswordTxRequest{
		Email:             req.Email,
		NewHashedPassword: hashedPassword,
		AfterChange: func(user db.User) error {
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10),
				asynq.Queue(worker.CriticaLQueue),
			}

			payload := &worker.PayloadSendVerificationEmail{
				Email: user.Email,
			}

			err := server.taskDistributor.TaskPasswordChangeVerificationEmail(c.Context(), payload, opts...)
			if err != nil {
				return fmt.Errorf("cannot send password changed email: %w", err)
			}
			return nil
		},
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := types.ChangePasswordResParams{
		Email:             userTx.User.Email,
		PasswordChangedAt: userTx.User.PasswordChangedAt,
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": response,
	})
}

func (s *Server) Health(ctx *fiber.Ctx) error {
	return ctx.SendString("OK")
}
