package app

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/routes"
	"github.com/gin-gonic/gin"
)

func (app *App) handlerV1(router *gin.RouterGroup) {
	divisionGroup := router.Group("/divisions")
	routes.DivisionRoutes(divisionGroup, app.delivery.division)
}
