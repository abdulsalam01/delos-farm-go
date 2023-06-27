package helper

import "github.com/go-playground/validator/v10"

var Validator *validator.Validate

func NewValidator() {
	// Init validator.
	Validator = validator.New()
}
