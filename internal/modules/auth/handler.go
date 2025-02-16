package auth

import (
	"time"

	"github.com/aayushchugh/timy-api/config/db"
	"github.com/aayushchugh/timy-api/config/env"
	"github.com/aayushchugh/timy-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func PostSignupHandler(c *fiber.Ctx) error {
	req := c.Locals("validatedBody").(*SignupRequestBody)

	var existingUser models.User
	if err := db.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "user already exists",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Error("error hashing password", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := db.DB.Create(&user).Error; err != nil {
		log.Error("error creating user", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully",
	})
}

func PostLoginHandler(c *fiber.Ctx) error {
	env := env.NewEnv()
	req := c.Locals("validatedBody").(*LoginRequestBody)

	// check if user exists
	var user models.User
	if err := db.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	// check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	// create access and refresh tokens

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(env.JWTSecret))

	if err != nil {
		log.Error("error generating access token", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 15).Unix(), // 15 days
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(env.JWTSecret))

	if err != nil {
		log.Error("error generating refresh token", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login successful",
		"payload": fiber.Map{
			"access_token":  accessTokenString,
			"refresh_token": refreshTokenString,
		},
	})
}

func GetMeHandler(c *fiber.Ctx) error {
	user_id := c.Locals("user_id")

	var user models.User
	if err := db.DB.Where("id = ?", user_id).First(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"payload": fiber.Map{
			"user": fiber.Map{
				"name":  user.Name,
				"email": user.Email,
			},
		},
	})
}
