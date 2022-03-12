package delivery

import (
	"time"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type UserDelivery interface {
	CreateUser(ctx *gin.Context)
}

func NewUserHandler(usecase usecase.UserUsecase) UserDelivery {
	return &userHandler{
		usecase: usecase,
	}
}

type userHandler struct {
	usecase usecase.UserUsecase
}

func (handler *userHandler) CreateUser(ctx *gin.Context) {
	res := handler.usecase.CreateUser(model.UserRegisterRequest{
		Name:      "tegar",
		Birthdate: time.Now(),
	})

	ctx.JSON(200, res)
}
