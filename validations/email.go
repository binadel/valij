package validations

import (
	"github.com/binadel/vali/constraints"
	"github.com/binadel/vali/validations/results"
)

type EmailValidation struct {
	stringValidation StringValidation
}

// Validate applies the validations constraints to the field value and returns the result.
func (v EmailValidation) Validate(value string) results.EmailResult {
	var result results.EmailResult
	result.StringResult = v.stringValidation.Validate(value)

	if len(result.Errors) == 0 {
		if email, err := constraints.ParseEmail(value); err != nil {
			result.Errors = append(result.Errors, err)
		} else {
			result.Email = email
		}
	}

	return result
}
