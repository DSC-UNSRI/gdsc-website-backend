package usecase

import (
	"context"
	"net/http"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/google/uuid"
)

func (usecase *divisionUsecaseImpl) GetDivision(req model.GetDivisionRequest) model.WebServiceResponse {
	uuid, err := uuid.Parse(req.ID)

	if err != nil {
		return model.WebServiceResponse{
			Message: "ID tidak valid",
			Status:  http.StatusUnprocessableEntity,
			Data:    nil,
		}
	}

	division, err := usecase.Store.GetDivision(context.Background(), uuid)
	if err != nil {
		return model.WebServiceResponse{
			Message: "Gagal mengambil data divisi",
			Status:  http.StatusInternalServerError,
			Data:    nil,
		}
	}

	return model.WebServiceResponse{
		Message: "Sukses mengambil data divisi",
		Status:  http.StatusOK,
		Data: map[string]interface{}{
			"division": division,
		},
	}
}
