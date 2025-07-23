package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeffrysusilo/go-wallet/services/wallet/controller"
	"github.com/jeffrysusilo/go-wallet/services/wallet/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Wallet service running 💰")
	})

	api := app.Group("/api", middleware.JWTProtected())
	api.Get("/balance", controller.GetBalance)
}
