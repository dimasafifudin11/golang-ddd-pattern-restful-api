package util

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func formatValidationErrors(err error) []ValidationError {
	var errors []ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		var message string
		switch err.Tag() {
		case "required":
			message = fmt.Sprintf("Kolom %s wajib diisi", err.Field())
		case "email":
			message = fmt.Sprintf("Kolom %s harus berupa format email yang valid", err.Field())
		case "min":
			message = fmt.Sprintf("Kolom %s harus memiliki panjang minimal %s karakter", err.Field(), err.Param())
		default:
			message = fmt.Sprintf("Kolom %s tidak valid", err.Field())
		}
		errors = append(errors, ValidationError{
			Field:   strings.ToLower(err.Field()),
			Message: message,
		})
	}
	return errors
}

func ValidateStruct(s interface{}) []ValidationError {
	err := validate.Struct(s)
	if err != nil {
		return formatValidationErrors(err)
	}
	return nil
}
