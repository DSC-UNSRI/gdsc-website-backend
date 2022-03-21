package delivery

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (handler *memberHandler) CreateMember(ctx *gin.Context) {
	var req model.CreateMemberRequest

	ok := utils.BindJSONAndValidate(ctx, &req)
	if !ok {
		return
	}

	res := handler.usecase.CreateMember(req)

	ctx.JSON(res.Status, res)
}
