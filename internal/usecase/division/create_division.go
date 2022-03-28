package usecase

import (
	"context"
	"net/http"

	postgresql "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/postgresql/sqlc"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (usecase *divisionUsecaseImpl) CreateDivision(req model.CreateDivisionRequest) model.WebServiceResponse {
	divisionDb, err := usecase.Store.CreateDivision(context.Background(), postgresql.CreateDivisionParams{
		Name: req.Name,
		Type: postgresql.DivisionType(req.Type),
	})
	if err != nil {
		return utils.ToWebServiceResponse("Gagal membuat divisi", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Divisi berhasil dibuat", http.StatusCreated, gin.H{
		"division": model.Division{
			ID:        divisionDb.ID,
			Name:      divisionDb.Name,
			Type:      string(divisionDb.Type),
			CreatedAt: divisionDb.CreatedAt,
		},
	})
}
