package usecase

import (
	"context"
	"errors"
	"net/http"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func (usecase *generationUsecaseImpl) GetGeneration(req model.GetOrDeleteGenerationRequest) model.WebServiceResponse {
	gen, err := usecase.Store.GetGeneration(context.Background(), uuid.MustParse(req.GenerationId))

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.WebServiceResponse{
				Message: "Data tidak ditemukan",
				Status:  http.StatusNotFound,
				Data:    nil,
			}
		}

		return model.WebServiceResponse{
			Message: "Terjadi kesalahan",
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	return model.WebServiceResponse{
		Message: "Berhasil mendapatkan data generasi",
		Status:  http.StatusOK,
		Data: gin.H{
			"generation": model.Generation(gen),
		},
	}
}
