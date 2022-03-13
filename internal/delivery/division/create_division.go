package delivery

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (handler *divisionHandler) CreateDivision(ctx *gin.Context) {
	var req model.CreateDivisionRequest

	ok := utils.BindJSONAndValidate(ctx, &req)
	if !ok {
		return
	}

	res := handler.usecase.CreateDivision(req)

	ctx.JSON(res.Status, res)
}
