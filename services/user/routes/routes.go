package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeffrysusilo/go-wallet/services/user/controller"
	"github.com/jeffrysusilo/go-wallet/services/user/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("User service running ✅")
	})

	// Public (untuk testing param)
	app.Get("/profile/:id", controller.GetUserProfileByID)

	// Protected
	api := app.Group("/api", middleware.JWTProtected())
	api.Get("/profile", controller.GetMyProfile)
}
