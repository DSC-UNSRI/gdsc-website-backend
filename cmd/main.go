package main

import (
	"context"
	"log"

	"github.com/DSC-UNSRI/gdsc-website-backend/config"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/app"
	"github.com/gin-gonic/gin/binding"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	config := config.New(".env")
	dbConfig, err := pgxpool.ParseConfig(config.PostgresDSN)
	if err != nil {
		log.Fatalf("Wrong dsn %v", err)
	}

	binding.Validator.Engine()

	dbPool, err := pgxpool.ConnectConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatalf("Can't connect to the database %v", err)
	}
	log.Println("Connected to postgres")
	app := app.New(config, dbPool)
	app.StartServer()
}
