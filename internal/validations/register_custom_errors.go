package validations

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func registerCustomErrors() {
	customErrors["required"] = func(field validator.FieldError, translatedFieldName string) string {
		return fmt.Sprintf("%s tidak boleh kosong", translatedFieldName)
	}

	customErrors["startswith"] = func(field validator.FieldError, translatedFieldName string) string {
		return fmt.Sprintf("%s harus berawalan %s", translatedFieldName, field.Value())
	}

	customErrors["uuid"] = func(field validator.FieldError, translatedFieldName string) string {
		return "ID invalid"
	}
}
