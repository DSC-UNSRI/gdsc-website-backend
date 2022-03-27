package usecase

import (
	"context"
	"net/http"

	postgresql "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/postgresql/sqlc"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (usecase *roleUsecaseImpl) ListRole(req model.ListRequest) model.WebServiceResponse {
	roles, err := usecase.Store.ListRole(context.Background(), postgresql.ListRoleParams{
		Limit:  int32(req.PageSize),
		Offset: int32((req.PageNumber - 1) * req.PageSize),
	})

	if err != nil {
		return utils.ToWebServiceResponse("Gagal mendapatkan data role", http.StatusInternalServerError, nil)
	}

	var newRoles []model.Role

	for _, role := range roles {
		newRoles = append(newRoles, model.Role(role))
	}

	return utils.ToWebServiceResponse("Sukses mendapatkan data role", http.StatusOK, gin.H{
		"roles": newRoles,
		"page_info": gin.H{
			"number":              req.PageNumber,
			"size":                req.PageSize,
			"current_page_length": len(roles),
		},
	})
}
