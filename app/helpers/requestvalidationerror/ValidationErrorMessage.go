package requestvalidationerror

import (
	"github.com/go-playground/validator/v10"
)

type ValidationField struct {
	Field   string
	Message string
}

func GetvalidationError(err error) []ValidationField {
	var validationFields []ValidationField
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, validationError := range ve {
			switch validationError.Tag() {
			case "required":
				validationFields = append(validationFields, ValidationField{validationError.Field(),
					validationError.Tag()})
			case "gte":
				validationFields = append(validationFields, ValidationField{validationError.Field(),
					"password length must be 6 or more"})
			case "min":
				validationFields = append(validationFields, ValidationField{validationError.Field(),
					"age must be 8 or more"})
			case "email":
				validationFields = append(validationFields, ValidationField{validationError.Field(),
					"format incorrect"})
			}
		}
	}

	return validationFields
}
