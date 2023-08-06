package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// NewValidator is a function that creates a new instance of the validator and registers a custom validation for UUID
func NewValidator() *validator.Validate {
	validate := validator.New()
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return false
		}
		return true
	})

	return validate
}

// ValidatorErrors is a function that takes a validation error and returns a map of fields and error messages
func ValidatorErrors(err error) map[string]string {
	fields := map[string]string{}
	
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}
