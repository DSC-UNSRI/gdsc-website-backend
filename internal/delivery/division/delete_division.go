package delivery

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (handler *divisionHandler) DeleteDivision(ctx *gin.Context) {
	var req model.DeleteDivisionRequest

	ok := utils.BindURIAndValidate(ctx, &req)

	if !ok {
		return
	}

	res := handler.usecase.DeleteDivision(req)
	ctx.JSON(res.Status, res)
}
