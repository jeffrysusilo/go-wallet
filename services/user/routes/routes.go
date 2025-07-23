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

	app.Get("/profile/:id", controller.GetUserProfileByID)

	api := app.Group("/api", middleware.JWTProtected())
	api.Get("/profile", controller.GetMyProfile)

	api.Put("/profile", controller.UpdateMyProfile)

	admin := api.Group("/admin", middleware.RoleGuard("admin"))
	admin.Get("/users", controller.GetAllUsers)

}
