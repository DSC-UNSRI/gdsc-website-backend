package delivery

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	defaultPageNumber = 1
	defaultPageSize   = 10
)

func (d *generationHandler) ListGenerations(ctx *gin.Context) {
	req := model.ListRequest{
		PageNumber: defaultPageNumber,
		PageSize:   defaultPageSize,
	}

	ok := utils.BindWith(ctx, &req, binding.Query)
	if !ok {
		return
	}

	res := d.usecase.ListGenerations(req)
	ctx.JSON(res.Status, res)
}
