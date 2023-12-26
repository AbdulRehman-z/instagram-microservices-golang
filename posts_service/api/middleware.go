package api

import (
	"fmt"
	"strings"

	"github.com/AbdulRehman-z/instagram-microservices/posts_service/token"
	"github.com/gofiber/fiber/v2"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.TokenVerifier) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("AuthMiddleware")
		authorizationHeader := c.GetReqHeaders()["Authorization"]
		if len(authorizationHeader) == 0 {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		fmt.Println("authorizationHeader: ", authorizationHeader)

		authorizationHeaderSplit := strings.Split(authorizationHeader[0], " ")
		if len(authorizationHeaderSplit) != 2 {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}
		fmt.Println("authorizationHeaderSplit: ", authorizationHeaderSplit)

		authorizationType := strings.ToLower(authorizationHeaderSplit[0])
		if authorizationType != authorizationTypeBearer {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}
		fmt.Println("authorizationType: ", authorizationType)

		accessToken := authorizationHeaderSplit[1]
		fmt.Println("accessToken: ", accessToken)
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return fiber.NewError(fiber.StatusForbidden, err.Error())
		}
		fmt.Println("payload: ", payload)

		c.Locals(authorizationPayloadKey, payload)
		return c.Next()
	}
}
