package validations

import (
	"github.com/binadel/vali/constraints"
	"github.com/binadel/vali/validations/results"
)

type RegexValidation struct {
	stringValidation StringValidation
}

// Validate applies the validations constraints to the field value and returns the result.
func (v RegexValidation) Validate(value string) results.RegexResult {
	var result results.RegexResult
	result.StringResult = v.stringValidation.Validate(value)

	if len(result.Errors) == 0 {
		if regex, err := constraints.ParseRegex(value); err != nil {
			result.Errors = append(result.Errors, err)
		} else {
			result.Regex = regex
		}
	}

	return result
}
