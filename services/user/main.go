package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/jeffrysusilo/go-wallet/services/user/config"
	"github.com/jeffrysusilo/go-wallet/services/user/routes"
)

func main() {
	godotenv.Load()
	config.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
