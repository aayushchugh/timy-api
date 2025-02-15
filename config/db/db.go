package db

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/aayushchugh/timy-api/config/env"
	"github.com/aayushchugh/timy-api/internal/models"
)

var DB *gorm.DB

func ConnectDB() {
	env := env.NewEnv()

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", env.DBHost, env.DBUser, env.DBPassword, env.DBName, env.DBPort)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	DB = db

	log.Info("Successfully connected to DB")
}
