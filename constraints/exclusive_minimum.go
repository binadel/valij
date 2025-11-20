package constraints

import (
	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var exclusiveMinimumError = errors.BasicError{
	Code:    core.RuleExclusiveMinimum,
	Message: "value must be greater than {{minimum}}",
}

// ExclusiveMinimumInt applies exclusive minimum constraint to an integer.
func ExclusiveMinimumInt(minimum int64) core.IntValidator {
	return func(value int64) core.Error {
		if value <= minimum {
			return &errors.IntParamError{
				BasicError: exclusiveMinimumError,
				ParamKey:   errors.ParamKeyMinimum,
				ParamValue: minimum,
			}
		}
		return nil
	}
}
