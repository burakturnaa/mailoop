package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func TrimValidator(fl validator.FieldLevel) bool {
	if fl.Field().Kind() == reflect.String {
		trimmed := strings.TrimSpace(fl.Field().String())

		field := fl.Field()
		if field.CanSet() {
			field.SetString(trimmed)
		}

		return true
	}
	return false
}
