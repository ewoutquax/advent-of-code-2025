package day02giftshop_services

import (
	validators "github.com/ewoutquax/advent-of-code-2025/internal/day-02-gift-shop/services/validator"
)

type FuncValidatorOption func(*validators.Validator)

func BuildValidator(opts ...FuncValidatorOption) validators.Validator {
	var defaultValidator = validators.Validator{ValidationType: validators.ValidationTypeBasic}

	for _, fnOpt := range opts {
		fnOpt(&defaultValidator)
	}

	return defaultValidator
}

func WithValidationTypeExtended() FuncValidatorOption {
	return func(v *validators.Validator) {
		v.ValidationType = validators.ValidationTypeExtended
	}
}
