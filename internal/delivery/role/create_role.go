package delivery

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (d *roleHandler) CreateRole(ctx *gin.Context) {
	var req model.CreateRoleRequest

	ok := utils.BindJSONAndValidate(ctx, &req)
	if !ok {
		return
	}

	res := d.usecase.CreateRole(req)
	ctx.JSON(res.Status, res)
}
