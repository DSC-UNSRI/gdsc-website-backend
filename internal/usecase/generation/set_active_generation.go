package usecase

import (
	"context"
	"errors"
	"net/http"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func (usecase *generationUsecaseImpl) SetActiveGeneration(req model.GetOrDeleteGenerationRequest) model.WebServiceResponse {
	var id uuid.NullUUID
	err := id.Scan(req.GenerationId)
	if err != nil {
		return utils.ToWebServiceResponse("ID tidak valid", http.StatusUnprocessableEntity, nil)
	}

	gen, err := usecase.Store.SetActiveGeneration(context.Background(), id)
	if err != nil {
		if errors.Is(pgx.ErrNoRows, err) {
			return utils.ToWebServiceResponse("Data generasi tidak ditemukan", http.StatusNotFound, nil)
		}

		return utils.ToWebServiceResponse("Terjadi kesalahan", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Sukses", http.StatusOK, gin.H{
		"generation": model.Generation(gen),
	})
}
