package usecase

import (
	"context"
	"net/http"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (usecase *roleUsecaseImpl) CreateRole(req model.CreateRoleRequest) model.WebServiceResponse {
	role, err := usecase.Store.CreateRole(context.Background(), req.Name)

	if err != nil {
		return utils.ToWebServiceResponse("Gagal membuat role", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Sukses membuat role", http.StatusCreated, gin.H{
		"role": model.Role(role),
	})
}
