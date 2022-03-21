package routes

import (
	delivery "github.com/DSC-UNSRI/gdsc-website-backend/internal/delivery/member"
	"github.com/gin-gonic/gin"
)

func MemberRoutes(router *gin.RouterGroup, delivery delivery.MemberDelivery) {
	router.POST("", delivery.CreateMember)
}
