package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v", err)
	}
	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping DB: %v", err)
	}
	log.Println("Connected to User DB")
}
