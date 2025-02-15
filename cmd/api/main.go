package main

import (
	"github.com/aayushchugh/timy-api/config/db"
	"github.com/aayushchugh/timy-api/config/env"
	"github.com/aayushchugh/timy-api/internal/handler/health"
	healthService "github.com/aayushchugh/timy-api/internal/service/health"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app := fiber.New()

	db.ConnectDB()

	hs := healthService.NewHealthService()
	healthHandler := health.NewHealthHandler(hs)

	env := env.NewEnv()
	port := "8000"

	if env.AppPort != "" {
		port = env.AppPort
	}

	health := app.Group("/health")
	health.Get("/", healthHandler.Check)

	log.Info("Server started on port 8000")
	app.Listen(":" + port)
}
