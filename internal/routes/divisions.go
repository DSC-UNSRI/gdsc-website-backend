package routes

import (
	delivery "github.com/DSC-UNSRI/gdsc-website-backend/internal/delivery/division"
	"github.com/gin-gonic/gin"
)

func DivisionRoutes(router *gin.RouterGroup, delivery delivery.DivisionDelivery) {
	router.POST("", delivery.CreateDivision)
	router.DELETE("/:id", delivery.DeleteDivision)
}
