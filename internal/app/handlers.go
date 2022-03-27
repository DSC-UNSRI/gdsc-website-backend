package app

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/routes"
	"github.com/gin-gonic/gin"
)

func (app *App) handlerV1(router *gin.RouterGroup) {
	divisionGroup := router.Group("/divisions")
	routes.DivisionRoutes(divisionGroup, app.delivery.division)

	generationGroup := router.Group("/generations")
	routes.GenerationRoutes(generationGroup, app.delivery.generation)

	roleGroup := router.Group("/roles")
	routes.RolesRoutes(roleGroup, app.delivery.role)
}
