package health

import (
	"time"

	"github.com/aayushchugh/timy-api/internal/domain/health"
)

type HealthService interface {
	Check() *health.Health
}

type healthService struct{}

func NewHealthService() HealthService {
	return &healthService{}
}

func (s *healthService) Check() *health.Health {
	return &health.Health{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
