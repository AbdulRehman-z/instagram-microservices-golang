package api

import (
	db "github.com/AbdulRehman-z/instagram-microservices/auth_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/types"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/util"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) VerifyEmail(c *fiber.Ctx) error {
	var req types.VerifyEmailReqParams
	err := c.QueryParser(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := util.CheckValidationErrors(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	verifyTxResult, err := s.store.VerifyEmailTx(c.Context(), db.VerifyEmailTxParams{
		EmailId:    int32(req.EmailId),
		SecretCode: req.SecretCode,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	resp := types.VerifyEmailResParams{
		IsEmailVerified: verifyTxResult.User.IsEmailVerified && verifyTxResult.VerifyEmail.IsUsed,
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": resp,
	})
}
