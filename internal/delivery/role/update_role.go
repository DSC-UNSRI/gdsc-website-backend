package delivery

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (d *roleHandler) UpdateRole(ctx *gin.Context) {
	var req model.UpdateRoleRequest

	ok := utils.BindURIAndValidate(ctx, &req.GetOrDeleteRoleRequest)
	if !ok {
		return
	}

	ok = utils.BindJSONAndValidate(ctx, &req.CreateRoleRequest)
	if !ok {
		return
	}

	res := d.usecase.UpdateRole(req)
	ctx.JSON(res.Status, res)
}
