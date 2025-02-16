package middlewares

import (
	"strings"

	"github.com/aayushchugh/timy-api/config/env"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserFromRequest(c *fiber.Ctx) error {
	env := env.NewEnv()

	// get jwt token from authorization header
	headers := c.GetReqHeaders()
	authorizationHeader := headers["Authorization"]

	if authorizationHeader == nil || len(authorizationHeader) == 0 {
		return c.Next()
	}

	token := strings.Split(authorizationHeader[0], " ")[1]

	// get user_id from token
	decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(env.JWTSecret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token expired",
		})
	}

	claims := decodedToken.Claims.(jwt.MapClaims)

	c.Locals("user_id", claims["user_id"])

	return c.Next()
}
