package config

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		AppPort: os.Getenv("APP_PORT"),
	}
}
