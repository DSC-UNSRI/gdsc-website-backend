package delivery

import (
	usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/member"
	"github.com/gin-gonic/gin"
)

type MemberDelivery interface {
	CreateMember(ctx *gin.Context)
}

var _ MemberDelivery = &memberHandler{}

func NewMemberDelivery(usecase usecase.MemberUsecase) MemberDelivery {
	return &memberHandler{
		usecase: usecase,
	}
}

type memberHandler struct {
	usecase usecase.MemberUsecase
}
