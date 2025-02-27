package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/suhailmuhammed157/simple_bank/utils"
)

var validateCurrency validator.Func = func(fl validator.FieldLevel) bool {
	currency, ok := fl.Field().Interface().(string)
	if ok {
		return utils.IsValidCurrency(currency)
	}
	return true
}
