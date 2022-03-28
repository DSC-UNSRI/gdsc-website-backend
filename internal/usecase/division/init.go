package usecase

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/db"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
)

//go:generate mockgen -source=./init.go -destination=./__mock__/division.go -package=mock_division_usecase
type DivisionUsecase interface {
	CreateDivision(model.CreateDivisionRequest) model.WebServiceResponse
	DeleteDivision(model.DeleteDivisionRequest) model.WebServiceResponse
	GetDivision(model.GetDivisionRequest) model.WebServiceResponse
}

var _ DivisionUsecase = &divisionUsecaseImpl{}

func NewDivisionUsecase(store db.Store) DivisionUsecase {
	return &divisionUsecaseImpl{
		Store: store,
	}
}

type divisionUsecaseImpl struct {
	db.Store
}
