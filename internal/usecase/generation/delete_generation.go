package usecase

import (
	"context"
	"net/http"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/google/uuid"
)

func (usecase *generationUsecaseImpl) DeleteGeneration(req model.GetOrDeleteGenerationRequest) model.WebServiceResponse {
	rows, err := usecase.Store.DeleteGeneration(context.Background(), uuid.MustParse(req.GenerationId))
	if err != nil {
		return model.WebServiceResponse{
			Message: "Gagal menghapus generasi",
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	} else if rows == 0 {
		return model.WebServiceResponse{
			Message: "Data generasi tidak ditemukan",
			Status:  http.StatusNotFound,
			Data:    nil,
		}
	}

	return model.WebServiceResponse{
		Message: "Berhasil menghapus data generasi",
		Status:  http.StatusOK,
		Data:    nil,
	}

}
