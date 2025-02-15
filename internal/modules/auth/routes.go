package auth

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	router := app.Group("/auth")

	router.Post("/signup", PostSignupHandler)
}
