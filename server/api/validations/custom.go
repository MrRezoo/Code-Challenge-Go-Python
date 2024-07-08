package validations

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

func GetValidationErrors(err error) *[]ValidationError {
	var validationErrors []ValidationError
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ValidationError{
				Property: err.Field(),
				Tag:      err.Tag(),
				Value:    err.Param(),
				Message:  err.Error(),
			})
		}
		return &validationErrors
	}
	return nil
}
