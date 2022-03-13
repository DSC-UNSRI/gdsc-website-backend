package validations

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var customErrors map[string]func(field string) string

var customFields map[string]string

func InitValidations(validate *validator.Validate) {
	customFields = map[string]string{}
	registerCustomFields()

	customErrors = map[string]func(field string) string{}
	registerCustomErrors()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}
