package usecase

import (
	"context"
	"net/http"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/google/uuid"
)

func (usecase *roleUsecaseImpl) DeleteRole(req model.GetOrDeleteRoleRequest) model.WebServiceResponse {
	rows, err := usecase.Store.DeleteRole(context.Background(), uuid.MustParse(req.RoleId))
	if err != nil {
		return utils.ToWebServiceResponse("Gagal menghapus role", http.StatusInternalServerError, nil)
	}

	if rows == 0 {
		return utils.ToWebServiceResponse("Role tidak ditemukan", http.StatusNotFound, nil)
	}

	return utils.ToWebServiceResponse("Role berhasil dihapus", http.StatusOK, nil)
}
