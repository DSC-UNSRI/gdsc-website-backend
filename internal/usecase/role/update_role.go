package usecase

import (
	"context"
	"net/http"

	postgresql "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/postgresql/sqlc"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (usecase *roleUsecaseImpl) UpdateRole(req model.UpdateRoleRequest) model.WebServiceResponse {
	role, err := usecase.Store.UpdateRole(context.Background(), postgresql.UpdateRoleParams{
		Name:   req.Name,
		Roleid: uuid.MustParse(req.RoleId),
	})

	if err != nil {
		return utils.ToWebServiceResponse("Gagal mengupdate role", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Sukses mengupdate role", http.StatusOK, gin.H{
		"role": model.Role(role),
	})
}
