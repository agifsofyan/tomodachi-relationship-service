package validator

import (
	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func Parse(err error) []FieldError {

	if err == nil {
		return nil
	}

	var errors []FieldError

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, FieldError{
			Field:   e.Field(),
			Message: message(e),
		})
	}

	return errors
}

func message(err validator.FieldError) string {

	switch err.Tag() {

	case "required":
		return "is required"

	case "email":
		return "must be a valid email"

	case "min":
		return "minimum length is " + err.Param()

	case "max":
		return "maximum length is " + err.Param()

	case "uuid":
		return "must be a valid UUID"

	default:
		return "is invalid"
	}
}
