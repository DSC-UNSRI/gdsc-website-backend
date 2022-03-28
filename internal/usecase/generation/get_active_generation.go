package usecase

import (
	"context"
	"errors"
	"net/http"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func (usecase *generationUsecaseImpl) GetActiveGeneration() model.WebServiceResponse {
	gen, err := usecase.Store.GetActiveGeneration(context.Background())
	if err != nil {
		if errors.Is(pgx.ErrNoRows, err) {
			return model.WebServiceResponse{
				Message: "Belum ada generasi yang aktif saat ini",
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
		Message: "Sukses mendapatkan data generasi",
		Status:  http.StatusOK,
		Data: gin.H{
			"active_generation": model.Generation(gen),
		},
	}
}
