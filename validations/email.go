package validations

import (
	"net/mail"

	"github.com/binadel/valij/constraints"
	"github.com/binadel/valij/validations/results"
)

type EmailValidation struct {
	stringValidation StringValidation
}

// Validate applies the validations constraints to the field value and returns the result.
func (v EmailValidation) Validate(value string) results.FormatResult[*mail.Address] {
	var result results.FormatResult[*mail.Address]
	result.StringResult = v.stringValidation.Validate(value)

	if len(result.Errors) == 0 {
		if email, err := constraints.ParseEmail(value); err != nil {
			result.Errors = append(result.Errors, err)
		} else {
			result.Format = email
		}
	}

	return result
}
