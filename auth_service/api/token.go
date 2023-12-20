package api

import (
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/types"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) renewAccessTokenHandler(c *fiber.Ctx) error {
	var req types.RenewAccessTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	refreshTokenPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}
	// get session from redis
	session, err := server.rStore.HGetAll(c.UserContext(), refreshTokenPayload.Email).Result()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if session["IsBlocked"] == "true" {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	if session["Email"] != refreshTokenPayload.Email {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(refreshTokenPayload.Email, refreshTokenPayload.UniqueId, server.Config.ACCESS_TOKEN_DURATION)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := types.RenewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": response,
	})
}
