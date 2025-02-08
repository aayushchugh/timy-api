package main

import (
	"github.com/aayushchugh/timy-api/config"
	"github.com/aayushchugh/timy-api/internal/handler/health"
	healthService "github.com/aayushchugh/timy-api/internal/service/health"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app := fiber.New()

	hs := healthService.NewHealthService()
	healthHandler := health.NewHealthHandler(hs)

	cfg := config.New()
	port := "8000"

	if cfg.AppPort != "" {
		port = cfg.AppPort
	}

	health := app.Group("/health")
	health.Get("/", healthHandler.Check)

	log.Info("Server started on port 8000")
	app.Listen(":" + port)
}
