package main

import (
	"github.com/aayushchugh/timy-api/config/db"
	"github.com/aayushchugh/timy-api/config/env"
	"github.com/aayushchugh/timy-api/internal/modules/auth"
	"github.com/aayushchugh/timy-api/internal/modules/health"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app := fiber.New()

	db.ConnectDB()

	env := env.NewEnv()
	port := "8000"

	if env.AppPort != "" {
		port = env.AppPort
	}

	health.SetupRoutes(app)
	auth.SetupRoutes(app)

	log.Info("Server started on port 8000")
	app.Listen(":" + port)
}
