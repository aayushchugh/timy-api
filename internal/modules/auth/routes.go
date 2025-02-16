package auth

import (
	"github.com/aayushchugh/timy-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	router := app.Group("/auth")

	router.Post("/signup", middlewares.ValidateRequestBody(func() interface{} {
		return &SignupRequestBody{}
	}), PostSignupHandler)

	router.Post("/login", middlewares.ValidateRequestBody(func() interface{} {
		return &LoginRequestBody{}
	}), PostLoginHandler)

	router.Get("/me", GetMeHandler)
}
