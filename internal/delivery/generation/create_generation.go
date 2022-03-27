package delivery

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (d *generationHandler) CreateGeneration(ctx *gin.Context) {
	var req model.CreateGenerationRequest

	ok := utils.BindJSONAndValidate(ctx, &req)
	if !ok {
		return
	}

	res := d.usecase.CreateGeneration(req)

	ctx.JSON(res.Status, res)
}
