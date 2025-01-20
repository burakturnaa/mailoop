package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func PhoneValidator(fl validator.FieldLevel) bool {
	phone := fl.Field().String()

	// e.g., +90 555 555 55 55 (Turkey format)
	phoneRegex := `^\+90 \d{3} \d{3} \d{2} \d{2}$`

	// Match the phone number against the regex
	match, _ := regexp.MatchString(phoneRegex, phone)
	return match
}
