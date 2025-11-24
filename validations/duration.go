package validations

import (
	"github.com/binadel/vali/constraints"
	"github.com/binadel/vali/validations/results"
)

type DurationValidation struct {
	stringValidation StringValidation
}

// Validate applies the validations constraints to the field value and returns the result.
func (v DurationValidation) Validate(value string) results.DurationResult {
	var result results.DurationResult
	result.StringResult = v.stringValidation.Validate(value)

	if len(result.Errors) == 0 {
		if duration, err := constraints.ParseDuration(value); err != nil {
			result.Errors = append(result.Errors, err)
		} else {
			result.Duration = duration
		}
	}

	return result
}
