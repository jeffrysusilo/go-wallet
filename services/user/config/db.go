package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")

	var err error
	for i := 0; i < 10; i++ {
		DB, err = pgxpool.New(context.Background(), dsn)
		if err == nil {
			err = DB.Ping(context.Background())
		}

		if err == nil {
			log.Println("✅ Connected to DB")
			return
		}

		log.Println("⏳ Waiting for DB... retrying in 2s")
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("❌ Unable to connect to DB: %v", err)
}
