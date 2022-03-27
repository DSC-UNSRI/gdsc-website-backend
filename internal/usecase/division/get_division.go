package usecase

import (
	"context"
	"net/http"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (usecase *divisionUsecaseImpl) GetDivision(req model.GetDivisionRequest) model.WebServiceResponse {
	uuid, err := uuid.Parse(req.ID)

	if err != nil {
		return utils.ToWebServiceResponse("ID tidak valid", http.StatusUnprocessableEntity, nil)
	}

	division, err := usecase.Store.GetDivision(context.Background(), uuid)
	if err != nil {
		return utils.ToWebServiceResponse("Gagal mengambil data divisi", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Sukses mengambil data divisi", http.StatusOK, gin.H{
		"division": division,
	})
}
