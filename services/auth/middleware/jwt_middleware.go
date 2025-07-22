package middleware

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Missing token"})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		// Simpan seluruh token ke Locals supaya bisa diakses RoleGuard
		c.Locals("user", token)

		// (opsional) Simpan user_id juga jika butuh di handler
		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user_id", claims["user_id"])

		return c.Next()
	}
}
