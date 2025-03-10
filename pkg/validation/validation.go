package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// validates all required empty fields
func ValidateEmptyStrings(fields ...string) error {
	for _, field := range fields {
		if err := validate.Var(field, "required"); err != nil {
			return errors.New("все поля должны быть заполнены")
		}
	}
	return nil
}

// validates positive numbers
func ValidatePositiveNumbers(nums ...any) error {
	for _, num := range nums {
		if err := validate.Var(num, "gt=0"); err != nil {
			return errors.New("число должно быть положительным")
		}
	}
	return nil
}

// validates fields of required struct
func ValidateStruct(s any) error {
	err := validate.Struct(s)
	if err != nil {
		return errors.New("поля содержат некорректные данные")
	}
	return nil
}
