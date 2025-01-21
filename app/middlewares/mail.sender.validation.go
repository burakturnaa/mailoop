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

// MailSenderValidation validates the request body against the provided struct
func MailSenderValidation(dtoType interface{}) fiber.Handler {
	// validator instance
	var validate = validator.New()

	return func(ctx *fiber.Ctx) error {

		// register custom validations
		validate.RegisterValidation("phone", utils.PhoneValidator)

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
				jsonFieldName := utils.GetJSONTag(dto, e.StructField())
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
