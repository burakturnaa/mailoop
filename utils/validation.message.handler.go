package utils

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// ValidationMessageHandler function to generate custom error messages
func ValidationMessageHandler(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email", e.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
	default:
		return fmt.Sprintf("%s is invalid", e.Field())
	}
}

func GetJSONTag(s interface{}, fieldName string) string {

	typ := reflect.TypeOf(s)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	// get the struct field by name
	field, found := typ.FieldByName(fieldName)
	if !found {
		return fieldName
	}
	// get the JSON tag
	jsonTag := field.Tag.Get("json")
	if jsonTag == "" || jsonTag == "-" {
		return fieldName
	}
	return jsonTag
}
