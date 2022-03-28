package delivery

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (d *generationHandler) DeleteGeneration(ctx *gin.Context) {
	var req model.GetOrDeleteGenerationRequest

	ok := utils.BindURIAndValidate(ctx, &req)
	if !ok {
		return
	}

	res := d.usecase.DeleteGeneration(req)
	ctx.JSON(res.Status, res)
}
