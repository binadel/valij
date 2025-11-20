package constraints

import (
	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var minimumError = errors.BasicError{
	Code:    core.RuleMinimum,
	Message: "value must be greater than or equal to {{minimum}}",
}

// MinimumInt applies minimum constraint to an integer.
func MinimumInt(minimum int64) core.IntValidator {
	return func(value int64) core.Error {
		if value < minimum {
			return &errors.IntParamError{
				BasicError: minimumError,
				ParamKey:   errors.ParamKeyMinimum,
				ParamValue: minimum,
			}
		}
		return nil
	}
}
