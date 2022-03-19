package usecase

import (
	"context"
	"net/http"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
)

func (usecase *divisionUsecaseImpl) CreateDivision(req model.CreateDivisionRequest) model.WebServiceResponse {
	divisionDb, err := usecase.Store.CreateDivision(context.Background(), req.Name)
	if err != nil {
		return utils.ToWebServiceResponse("Gagal membuat divisi", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Divisi berhasil dibuat", http.StatusCreated, map[string]interface{}{
		"division": model.Division(divisionDb),
	})
}
