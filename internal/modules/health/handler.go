package health

import "github.com/gofiber/fiber/v2"

func GetHealthHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "OK",
	})
}
