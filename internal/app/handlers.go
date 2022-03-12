package app

import (
	"github.com/gin-gonic/gin"
)

func (app *App) createHandlers() *gin.Engine {
	router := gin.Default()
	router.GET("/", app.delivery.user.CreateUser)

	return router
}
