package metadata

import validator "gopkg.in/go-playground/validator.v9"

var govalidator = createValidator()

// InputValidation defines an interface for all "input submission" structs used for deserialization
type InputValidation interface {
	Validate() error
}

func createValidator() *validator.Validate {
	var validator = validator.New()

	return validator
}
