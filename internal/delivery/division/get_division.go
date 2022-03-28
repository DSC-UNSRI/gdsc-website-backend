package delivery

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (d *divisionHandler) GetDivision(ctx *gin.Context) {
	var req model.GetDivisionRequest

	ok := utils.BindURIAndValidate(ctx, &req)
	if !ok {
		return
	}

	res := d.usecase.GetDivision(req)

	ctx.JSON(res.Status, res)
}
