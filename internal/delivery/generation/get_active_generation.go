package delivery

import "github.com/gin-gonic/gin"

func (d *generationHandler) GetActiveGeneration(ctx *gin.Context) {
	res := d.usecase.GetActiveGeneration()

	ctx.JSON(res.Status, res)
}
