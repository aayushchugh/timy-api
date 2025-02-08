package health

import (
	"github.com/aayushchugh/timy-api/internal/service/health"
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
	healthService health.HealthService
}

func NewHealthHandler(healthService health.HealthService) *HealthHandler {
	return &HealthHandler{
		healthService: healthService,
	}
}

func (h *HealthHandler) Check(c *fiber.Ctx) error {
	return c.JSON(h.healthService.Check())
}
