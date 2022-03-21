package usecase

import (
	"context"
	"net/http"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/google/uuid"
)

func (usecase *divisionUsecaseImpl) DeleteDivision(req model.DeleteDivisionRequest) model.WebServiceResponse {
	rowsAffected, _ := usecase.Store.DeleteDivision(context.Background(), uuid.MustParse(req.ID))
	if rowsAffected == 0 {
		return utils.ToWebServiceResponse("Gagal menghapus divisi", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Divisi berhasil dihapus", http.StatusOK, nil)
}
