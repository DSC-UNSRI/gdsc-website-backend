package usecase

import (
	"context"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/db"
	postgresql "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/postgresql/sqlc"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
)

type UserUsecase interface {
	CreateUser(model.UserRegisterRequest) model.WebServiceResponse
}

func NewUserUsecase(store *db.Store) UserUsecase {
	return &userUsecaseImpl{
		Store: store,
	}
}

var _ UserUsecase = &userUsecaseImpl{}

type userUsecaseImpl struct {
	*db.Store
}

func (usecase *userUsecaseImpl) CreateUser(req model.UserRegisterRequest) model.WebServiceResponse {
	user, err := usecase.Store.CreateUser(context.Background(), postgresql.CreateUserParams{
		Name:      req.Name,
		Birthdate: req.Birthdate,
	})

	if err != nil {
		return model.WebServiceResponse{
			Message: "Error bos",
			Status:  500,
			Data:    nil,
		}
	}

	return model.WebServiceResponse{
		Message: "ini ga error bos",
		Status:  200,
		Data: model.UserRegisterResponse{
			Name:      user.Name,
			Birthdate: user.Birthdate,
		},
	}
}
