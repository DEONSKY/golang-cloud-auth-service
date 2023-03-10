package server

import (
	"github.com/forfam/authentication-service/customerror"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct[T any](data T) error {
	var errors []customerror.ValidationErrorData
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element customerror.ValidationErrorData
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, element)
		}
		return &customerror.ValidationErrors{Errors: errors}
	}
	return nil
}
