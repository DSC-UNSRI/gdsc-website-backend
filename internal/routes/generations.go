package routes

import (
	delivery "github.com/DSC-UNSRI/gdsc-website-backend/internal/delivery/generation"
	"github.com/gin-gonic/gin"
)

func GenerationRoutes(router *gin.RouterGroup, delivery delivery.GenerationDelivery) {
	router.GET("", delivery.ListGenerations)
	router.POST("", delivery.CreateGeneration)

	router.GET("/active", delivery.GetActiveGeneration)
	router.GET("/:id", delivery.GetGeneration)
	router.PUT("/:id", delivery.UpdateGeneration)
	router.DELETE("/:id", delivery.DeleteGeneration)
	
	router.PATCH("/:id/set-active", delivery.SetActiveGeneration)
}
