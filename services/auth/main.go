package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/jeffrysusilo/go-wallet/services/auth/config"
	"github.com/jeffrysusilo/go-wallet/services/auth/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app) 

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
