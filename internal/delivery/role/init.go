package delivery

import (
	usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/role"
	"github.com/gin-gonic/gin"
)

type RoleDelivery interface {
	CreateRole(ctx *gin.Context)
	DeleteRole(ctx *gin.Context)
	GetRole(ctx *gin.Context)
	ListRoles(ctx *gin.Context)
	UpdateRole(ctx *gin.Context)
}

var _ RoleDelivery = &roleHandler{}

func NewRoleDelivery(usecase usecase.RoleUsecase) RoleDelivery {
	return &roleHandler{
		usecase: usecase,
	}
}

type roleHandler struct {
	usecase usecase.RoleUsecase
}
