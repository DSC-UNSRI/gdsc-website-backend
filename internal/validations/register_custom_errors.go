package validations

import "fmt"

func registerCustomErrors() {
	customErrors["required"] = func(field string) string {
		var name string
		val, ok := customFields[field]
		if ok {
			name = val
		} else {
			name = field
		}
		return fmt.Sprintf("%s tidak boleh kosong", name)
	}
}
