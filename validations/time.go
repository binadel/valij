package validations

import (
	"time"

	"github.com/binadel/valij/constraints"
	"github.com/binadel/valij/validations/results"
)

type TimeValidation struct {
	stringValidation StringValidation
	onlyTime         bool
	onlyDate         bool
}

// Validate applies the validations constraints to the field value and returns the result.
func (v TimeValidation) Validate(value string) results.FormatResult[time.Time] {
	var result results.FormatResult[time.Time]
	result.StringResult = v.stringValidation.Validate(value)

	if len(result.Errors) == 0 {
		if v.onlyTime {
			if t, err := constraints.ParseTime(value); err != nil {
				result.Errors = append(result.Errors, err)
			} else {
				result.Format = t
			}
		} else if v.onlyDate {
			if t, err := constraints.ParseDate(value); err != nil {
				result.Errors = append(result.Errors, err)
			} else {
				result.Format = t
			}
		} else {
			if t, err := constraints.ParseDateTime(value); err != nil {
				result.Errors = append(result.Errors, err)
			} else {
				result.Format = t
			}
		}
	}

	return result
}
