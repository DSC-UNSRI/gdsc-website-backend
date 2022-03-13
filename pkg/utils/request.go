package utils

import (
	"encoding/json"
	"net/http"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindJSONAndValidate(ctx *gin.Context, i interface{}) bool {
	err := ctx.ShouldBindJSON(i)
	if err == nil {
		return true
	}

	if errValidations, ok := err.(validator.ValidationErrors); ok {
		res := validations.HandleValidationErrors(errValidations)
		ctx.JSON(res.Status, res)
	} else if _, ok := err.(*json.UnmarshalTypeError); ok {
		ctx.JSON(http.StatusBadRequest, model.WebServiceResponse{
			Message: "Schema request tidak valid",
			Status:  http.StatusBadRequest,
			Data:    nil,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, ToWebServiceResponse("Internal server error", http.StatusInternalServerError, nil))
	}

	return false
}