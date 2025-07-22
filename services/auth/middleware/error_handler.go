package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	log.Println("🔥 ERROR:", err)
	return c.Status(500).JSON(fiber.Map{
		"error": err.Error(),
	})
}
