package models

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResp struct {
	Field string
	Tag   string
	Value string
}

func validateModelsForErrors(err error) []*ErrorResp {
	var errors []*ErrorResp
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResp
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
