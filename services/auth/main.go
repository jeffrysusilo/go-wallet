package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/jeffrysusilo/go-wallet/services/auth/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB() // <-- inisialisasi koneksi DB

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Auth Service running 🚀")
	})

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
