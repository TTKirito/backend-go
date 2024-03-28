package api

import (
	"github.com/TTKirito/backend-go/utils"
	"github.com/go-playground/validator/v10"
)

var validStatus validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if status, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.IsSupportedStatus((status))
	}

	return false
}

var validPosition validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if position, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.IsSupportedPosition(position)
	}
	return false
}

var validGender validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if gender, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.IsSupportedGender(gender)
	}
	return false
}
