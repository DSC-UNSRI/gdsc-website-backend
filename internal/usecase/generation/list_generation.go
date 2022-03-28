package usecase

import (
	"context"
	"net/http"

	postgresql "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/postgresql/sqlc"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (usecase *generationUsecaseImpl) ListGenerations(req model.ListRequest) model.WebServiceResponse {
	param := postgresql.ListGenerationParams{
		Limit:  int32(req.PageSize),
		Offset: int32(req.PageSize * (req.PageNumber - 1)),
	}
	gens, err := usecase.Store.ListGeneration(context.Background(), param)
	if err != nil {
		return model.WebServiceResponse{
			Message: "Terjadi kesalahan",
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	var newGens []model.Generation
	for _, gen := range gens {
		newGens = append(newGens, model.Generation(gen))
	}

	return model.WebServiceResponse{
		Message: "Sukses",
		Status:  http.StatusOK,
		Data: gin.H{
			"generations": newGens,
			"page_info": gin.H{
				"number":              req.PageNumber,
				"size":                req.PageSize,
				"current_page_length": len(gens),
			},
		},
	}
}
