package api

import (
	"strings"

	"github.com/AbdulRehman-z/instagram-microservices/posts_service/token"
	"github.com/gofiber/fiber/v2"
)

var (
	authorizationBearerType = "Bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenVerifier token.TokenVerifier) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorizationHeader := c.Get(fiber.HeaderAuthorization)
		if len(authorizationHeader) == 0 {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		authorizationHeaderSplit := strings.Split(authorizationHeader, " ")
		if len(authorizationHeaderSplit) != 2 {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		authorizationType := strings.ToLower(authorizationHeaderSplit[0])
		if authorizationType != authorizationBearerType {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		accessToken := authorizationHeaderSplit[1]
		payload, err := tokenVerifier.VerifyToken(accessToken)
		if err != nil {
			return fiber.NewError(fiber.StatusForbidden, err.Error())
		}

		c.Locals(authorizationPayloadKey, payload)
		return c.Next()
	}
}
