package api

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/AbdulRehman-z/instagram-microservices/auth_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/types"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/util"
	"github.com/gofiber/fiber/v2"
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

	args := db.RegisterUserParams{
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	user, err := server.store.RegisterUser(context.Background(), args)
	if err != nil {
		if pqerr, ok := err.(*pq.Error); ok {
			switch pqerr.Code.Name() {
			case "unique_violation", "foreign_key_violation":
				return fiber.NewError(fiber.StatusBadRequest, "user already registered")
			}
		}
	}

	response := types.RegisterUserResParams{
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
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
