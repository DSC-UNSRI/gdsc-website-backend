package usecase

import (
	"context"
	"fmt"
	"net/http"

	postgresql "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/postgresql/sqlc"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

func (usecase *generationUsecaseImpl) UpdateGeneration(req model.UpdateGenerationRequest) model.WebServiceResponse {
	gen, err := usecase.Store.UpdateGeneration(context.Background(), postgresql.UpdateGenerationParams{
		Year:         req.Year,
		Generationid: uuid.MustParse(req.GenerationId),
	})
	if err != nil {
		if err, ok := err.(*pgconn.PgError); ok && err.Code == pgerrcode.UniqueViolation {
			return utils.ToWebServiceResponse(fmt.Sprintf("Tahun %s sudah ada", req.Year), http.StatusUnprocessableEntity, nil)
		}

		return utils.ToWebServiceResponse("Gagal memperbarui data generasi", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Berhasil memperbarui data generasi", http.StatusOK, gin.H{
		"generation": model.Generation(gen),
	})
}
