package main

import (
	"github.com/aayushchugh/timy-api/config/db"
	"github.com/aayushchugh/timy-api/config/env"
	"github.com/aayushchugh/timy-api/internal/middlewares"
	"github.com/aayushchugh/timy-api/internal/modules/auth"
	"github.com/aayushchugh/timy-api/internal/modules/health"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	db.ConnectDB()

	env := env.NewEnv()
	port := "8000"

	if env.AppPort != "" {
		port = env.AppPort
	}

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${method} ${path}",
		TimeFormat: "15:04:05",
		TimeZone:   "Local",
	}))
	app.Use(middlewares.GetUserFromRequest)
	health.SetupRoutes(app)
	auth.SetupRoutes(app)

	log.Info("Server started on port 8000")
	app.Listen(":" + port)
}
