package transfers

import "github.com/go-playground/validator/v10"

// TODO: Not working as of now lmaooo
var ValidCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
    return IsSupportedCurrency(currency)
	}
  return false
}
