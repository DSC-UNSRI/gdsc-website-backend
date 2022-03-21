package app

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/routes"
	"github.com/gin-gonic/gin"
)

func (app *App) handlerV1(router *gin.RouterGroup) {
	divisionGroup := router.Group("/divisions")
	memberGroup := router.Group("/members")
	routes.DivisionRoutes(divisionGroup, app.delivery.division)
	routes.MemberRoutes(memberGroup, app.delivery.member)
}
