package usecase

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/db"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
)

//go:generate mockgen -source=./init.go -destination=./__mock__/generation.go -package=mock_generation_usecase
type RoleUsecase interface {
	CreateRole(model.CreateRoleRequest) model.WebServiceResponse
	UpdateRole(model.UpdateRoleRequest) model.WebServiceResponse
	DeleteRole(model.GetOrDeleteRoleRequest) model.WebServiceResponse
	GetRole(model.GetOrDeleteRoleRequest) model.WebServiceResponse
	ListRole(model.ListRequest) model.WebServiceResponse
}

var _ RoleUsecase = &roleUsecaseImpl{}

func NewGenerationUsecase(store db.Store) RoleUsecase {
	return &roleUsecaseImpl{
		Store: store,
	}
}

type roleUsecaseImpl struct {
	db.Store
}
