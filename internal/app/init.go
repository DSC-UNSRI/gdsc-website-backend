package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/DSC-UNSRI/gdsc-website-backend/config"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/db"
	"github.com/jackc/pgx/v4/pgxpool"
)

type App struct {
	Config   config.Config
	delivery deliveries
	usecase  usecases
	store    *db.Store
}

func New(config config.Config, dbPool *pgxpool.Pool) App {
	var app App
	app.Config = config
	app.store = db.NewStore(dbPool)
	app.initUsecase()
	app.initDelivery()
	return app
}

func (app *App) StartServer() {
	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	router := app.createHandlers()
	address := fmt.Sprintf("%s:%s", app.Config.AppHost, app.Config.AppPort)
	log.Printf("Server listening on %v\n", address)

	srv := &http.Server{
		Addr:    address,
		Handler: router,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatalf("Cannot start server %v\n", err)
		}
	}()

	<-osSignalChan
	err := srv.Close()
	if err != nil {
		log.Fatalf("cannot shutdown server %v", err)
	}
	log.Println("1 (3 h\nServer exiting")
}
