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

func (usecase *roleUsecaseImpl) GetRole(req model.GetOrDeleteRoleRequest) model.WebServiceResponse {
	role, err := usecase.Store.GetRole(context.Background(), uuid.MustParse(req.RoleId))

	if err != nil {
		if errors.Is(pgx.ErrNoRows, err) {
			return utils.ToWebServiceResponse("Role tidak ditemukan", http.StatusNotFound, nil)
		}

		return utils.ToWebServiceResponse("Gagal mendapatkan data role", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Sukses mendapatkan data role", http.StatusOK, gin.H{
		"role": model.Role(role),
	})
}
