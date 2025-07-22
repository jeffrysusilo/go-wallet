package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeffrysusilo/go-wallet/services/auth/controller"
	"github.com/jeffrysusilo/go-wallet/services/auth/middleware"
)


func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Auth Service running 🚀")
	})

	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)

	app.Get("/me", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		return c.JSON(fiber.Map{"user_id": userID})
	})
}