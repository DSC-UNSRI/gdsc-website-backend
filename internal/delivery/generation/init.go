package delivery

import (
	usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/generation"
	"github.com/gin-gonic/gin"
)

type GenerationDelivery interface {
	CreateGeneration(ctx *gin.Context)
	DeleteGeneration(ctx *gin.Context)
	GetGeneration(ctx *gin.Context)
	ListGenerations(ctx *gin.Context)
	UpdateGeneration(ctx *gin.Context)
	GetActiveGeneration(ctx *gin.Context)
	SetActiveGeneration(ctx *gin.Context)
}

var _ GenerationDelivery = &generationHandler{}

func NewGenerationDelivery(usecase usecase.GenerationUsecase) GenerationDelivery {
	return &generationHandler{
		usecase: usecase,
	}
}

type generationHandler struct {
	usecase usecase.GenerationUsecase
}
