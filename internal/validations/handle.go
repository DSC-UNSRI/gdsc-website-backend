package validations

import (
	"fmt"
	"net/http"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/go-playground/validator/v10"
)

func debugFieldError(v validator.FieldError) {
	fmt.Println("ini field", v.Field())
	fmt.Println("ini kind", v.Kind())
	fmt.Println("ini actual tag", v.ActualTag())
	fmt.Println("ini tag", v.Tag())
	fmt.Println("ini value", v.Value())
	fmt.Println("ini type", v.Type())
	fmt.Println("ini namespace", v.Namespace())
	fmt.Println("ini param", v.Param())
	fmt.Println("ini struct field", v.StructField())
	fmt.Println("ini struct namespace", v.StructNamespace())
}

func HandleValidationErrors(errs validator.ValidationErrors) (res model.ValidationErrorWebServiceResponse) {
	res.Status = http.StatusUnprocessableEntity
	var finalErrors []string
	for _, v := range errs {
		// debugFieldError(v)
		finalErrors = append(
			finalErrors,
			customErrors[v.Tag()](v.Field()),
		)
	}
	if len(finalErrors) > 0 {
		res.Message = finalErrors[0]
	}

	res.DetailErrors = finalErrors
	res.Data = nil
	return res
}
