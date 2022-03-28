package routes

import (
	delivery "github.com/DSC-UNSRI/gdsc-website-backend/internal/delivery/role"
	"github.com/gin-gonic/gin"
)

func RolesRoutes(router *gin.RouterGroup, delivery delivery.RoleDelivery) {
	router.POST("", delivery.CreateRole)
	router.GET("", delivery.ListRoles)

	router.GET("/:id", delivery.GetRole)
	router.PUT("/:id", delivery.UpdateRole)
	router.DELETE("/:id", delivery.DeleteRole)
}
