package constraints

import (
	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var exclusiveMaximumError = errors.BasicError{
	Code:    core.RuleExclusiveMaximum,
	Message: "value must be less than {{maximum}}",
}

// ExclusiveMaximumInt applies exclusive maximum constraint to an integer.
func ExclusiveMaximumInt(maximum int64) core.IntValidator {
	return func(value int64) core.Error {
		if value >= maximum {
			return &errors.IntParamError{
				BasicError: exclusiveMaximumError,
				ParamKey:   errors.ParamKeyMaximum,
				ParamValue: maximum,
			}
		}
		return nil
	}
}
