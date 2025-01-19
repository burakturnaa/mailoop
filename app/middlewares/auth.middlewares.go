package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/burakturnaa/mailoop.git/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// validator instance
var validate = validator.New()

// AuthMiddleware validates the request body against the provided struct
func AuthMiddleware(dtoType interface{}) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// rebuild DTO for every request
		dto := reflect.New(reflect.TypeOf(dtoType).Elem()).Interface()

		if len(ctx.Body()) == 0 {
			response := utils.BuildResponse(4001, "Invalid request body format", nil, nil)
			return ctx.Status(http.StatusBadRequest).JSON(response)
		}

		errorMessages := make(map[string]string)

		// parsing the body
		if err := ctx.BodyParser(dto); err != nil {
			var unmarshalTypeError *json.UnmarshalTypeError
			// check if the error is a JSON unmarshalling type error
			if errors.As(err, &unmarshalTypeError) {
				fieldName := unmarshalTypeError.Field
				errorMessages[fieldName] = fmt.Sprintf("%s must be a %s", fieldName, unmarshalTypeError.Type.String())
			} else {
				// handle other parsing errors
				response := utils.BuildResponse(4001, "Invalid request body format", err.Error(), nil)
				return ctx.Status(http.StatusBadRequest).JSON(response)
			}
		}

		// validate the parsed struct
		if err := validate.Struct(dto); err != nil {
			validationErrors := err.(validator.ValidationErrors)

			for _, e := range validationErrors {
				jsonFieldName := getJSONTag(dto, e.StructField())
				errorMessages[jsonFieldName] = utils.ValidationMessageHandler(e)
			}
		}

		if len(errorMessages) > 0 {
			response := utils.BuildResponse(4002, "Validation error", errorMessages, nil)
			return ctx.Status(http.StatusBadRequest).JSON(response)
		}

		// next middleware/handler
		return ctx.Next()
	}
}

func getJSONTag(s interface{}, fieldName string) string {

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
