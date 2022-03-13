package delivery

import (
	usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/division"
	"github.com/gin-gonic/gin"
)

type DivisionDelivery interface {
	CreateDivision(ctx *gin.Context)
}

var _ DivisionDelivery = &divisionHandler{}

func NewDivisionDelivery(usecase usecase.DivisionUsecase) DivisionDelivery {
	return &divisionHandler{
		usecase: usecase,
	}
}

type divisionHandler struct {
	usecase usecase.DivisionUsecase
}