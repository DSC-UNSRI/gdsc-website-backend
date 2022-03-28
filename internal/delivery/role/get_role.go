package delivery

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (d *roleHandler) GetRole(ctx *gin.Context) {
	var req model.GetOrDeleteRoleRequest

	ok := utils.BindURIAndValidate(ctx, &req)
	if !ok {
		return
	}

	res := d.usecase.GetRole(req)
	ctx.JSON(res.Status, res)
}
