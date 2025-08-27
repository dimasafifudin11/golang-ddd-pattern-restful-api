package util

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		// This can be enhanced to return more specific error messages
		return errors.New("validation failed: " + err.Error())
	}
	return nil
}
