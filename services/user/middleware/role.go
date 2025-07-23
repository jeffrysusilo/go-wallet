package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RoleGuard(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user")
		if user == nil {
			return c.Status(403).JSON(fiber.Map{"error": "Access denied"})
		}

		claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
		role := claims["role"]
		if role != requiredRole {
			return c.Status(403).JSON(fiber.Map{"error": "Insufficient permissions"})
		}

		return c.Next()
	}
}
