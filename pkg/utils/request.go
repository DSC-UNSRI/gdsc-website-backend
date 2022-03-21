package utils

import (
	"encoding/json"
	"fmt"
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
	validate(ctx, err)
	return false
}

func BindURIAndValidate(ctx *gin.Context, i interface{}) bool {

	err := ctx.ShouldBindUri(i)
	fmt.Println("Bind")
	if err == nil {
		return true
	}
	validate(ctx, err)
	return false
}

func validate(ctx *gin.Context, err error) {
	fmt.Println("Validate")
	if errValidations, ok := err.(validator.ValidationErrors); ok {
		res := validations.HandleValidationErrors(errValidations)
		fmt.Println("Validate 1")
		ctx.JSON(res.Status, res)
	} else if _, ok := err.(*json.UnmarshalTypeError); ok {
		fmt.Println("Validate 2")
		ctx.JSON(http.StatusBadRequest, model.WebServiceResponse{
			Message: "Schema request tidak valid",
			Status:  http.StatusBadRequest,
			Data:    nil,
		})
	} else {
		fmt.Println("Validate 3")
		ctx.JSON(http.StatusInternalServerError, ToWebServiceResponse("Internal server error", http.StatusInternalServerError, nil))
	}
}
