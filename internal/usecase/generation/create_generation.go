package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

func (usecase *generationUsecaseImpl) CreateGeneration(req model.CreateGenerationRequest) model.WebServiceResponse {
	gen, err := usecase.Store.CreateGeneration(context.Background(), req.Year)
	if err != nil {
		if err, ok := err.(*pgconn.PgError); ok && err.Code == pgerrcode.UniqueViolation {
			return utils.ToWebServiceResponse(fmt.Sprintf("Tahun %s sudah ada", req.Year), http.StatusUnprocessableEntity, nil)
		}

		return utils.ToWebServiceResponse("Gagal membuat generasi baru", http.StatusInternalServerError, nil)
	}

	return model.WebServiceResponse{
		Message: "Berhasil membuat generasi baru",
		Status:  http.StatusCreated,
		Data: map[string]interface{}{
			"generation": model.Generation(gen),
		},
	}

}
