package usecase

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/db"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
)

//go:generate mockgen -source=./init.go -destination=./__mock__/generation.go -package=mock_generation_usecase
type GenerationUsecase interface {
	CreateGeneration(model.CreateGenerationRequest) model.WebServiceResponse
	UpdateGeneration(model.UpdateGenerationRequest) model.WebServiceResponse
	DeleteGeneration(model.GetOrDeleteGenerationRequest) model.WebServiceResponse
	GetActiveGeneration() model.WebServiceResponse
	SetActiveGeneration(model.GetOrDeleteGenerationRequest) model.WebServiceResponse
	ListGenerations(model.ListRequest) model.WebServiceResponse
	GetGeneration(model.GetOrDeleteGenerationRequest) model.WebServiceResponse
}

var _ GenerationUsecase = &generationUsecaseImpl{}

func NewGenerationUsecase(store db.Store) GenerationUsecase {
	return &generationUsecaseImpl{
		Store: store,
	}
}

type generationUsecaseImpl struct {
	db.Store
}
