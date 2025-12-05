package validations

import (
	"regexp"

	"github.com/binadel/valij/constraints"
	"github.com/binadel/valij/validations/results"
)

type RegexValidation struct {
	stringValidation StringValidation
}

// Validate applies the validations constraints to the field value and returns the result.
func (v RegexValidation) Validate(value string) results.FormatResult[*regexp.Regexp] {
	var result results.FormatResult[*regexp.Regexp]
	result.StringResult = v.stringValidation.Validate(value)

	if len(result.Errors) == 0 {
		if regex, err := constraints.ParseRegex(value); err != nil {
			result.Errors = append(result.Errors, err)
		} else {
			result.Format = regex
		}
	}

	return result
}
