package validator

import "github.com/go-playground/validator"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate(s any) error {
	return validate.Struct(s)
}

func GetValidator() *validator.Validate {
	return validate
}
