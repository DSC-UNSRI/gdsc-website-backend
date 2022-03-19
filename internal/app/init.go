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
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/validations"
	"github.com/TwiN/go-color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
)

type App struct {
	Config   config.Config
	delivery deliveries
	usecase  usecases
	store    db.Store
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
	if app.Config.Env == config.EnvProd {
		fmt.Println(
			color.Ize(color.Yellow, color.InBold("\nAPP RUN IN PRODUCTION MODE\n")),
		)
	} else {
		fmt.Println(
			color.Ize(color.Red, color.InBold("\nAPP RUN IN DEVELOPMENT MODE\n")),
		)
	}

	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validations.InitValidations(validator)
	}
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
	fmt.Println()
	log.Println("Server exiting")
}

func (app *App) createHandlers() *gin.Engine {
	router := gin.Default()

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowHeaders = append(corsCfg.AllowHeaders, "Accept")
	if app.Config.Env == config.EnvProd {
		corsCfg.AllowAllOrigins = false
		corsCfg.AllowOrigins = []string{app.Config.AllowedOrigin}
	} else {
		corsCfg.AllowAllOrigins = true
	}

	router.Use(cors.New(corsCfg))

	v1 := router.Group("/api/v1/")
	app.handlerV1(v1)

	routes := router.Routes()
	if gin.Mode() == gin.DebugMode {
		fmt.Println()
		for _, v := range routes {
			path := color.InBold(v.Path)
			method := color.InYellow(fmt.Sprintf("%-6s", v.Method))
			fmt.Println(method, path)
		}
		fmt.Println()
	}
	return router
}
