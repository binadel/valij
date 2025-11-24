package validations

import (
	"github.com/binadel/vali/constraints"
	"github.com/binadel/vali/validations/results"
)

type UriValidation struct {
	stringValidation StringValidation
}

// Validate applies the validations constraints to the field value and returns the result.
func (v UriValidation) Validate(value string) results.UriResult {
	var result results.UriResult
	result.StringResult = v.stringValidation.Validate(value)

	if len(result.Errors) == 0 {
		if uri, err := constraints.ParseUri(value); err != nil {
			result.Errors = append(result.Errors, err)
		} else {
			result.URI = uri
		}
	}

	return result
}
