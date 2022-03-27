package delivery

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (d *generationHandler) UpdateGeneration(ctx *gin.Context) {
	var req model.UpdateGenerationRequest

	ok := utils.BindURIAndValidate(ctx, &req.GetOrDeleteGenerationRequest)
	if !ok {
		return
	}

	ok = utils.BindJSONAndValidate(ctx, &req.CreateGenerationRequest)
	if !ok {
		return
	}

	res := d.usecase.UpdateGeneration(req)
	ctx.JSON(res.Status, res)
}
