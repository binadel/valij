package constraints

import (
	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var multipleOfError = errors.BasicError{
	Code:    core.RuleMultipleOf,
	Message: "value must be a multiple of {{factor}}",
}

// MultipleOfInt applies multiple constraint to an integer.
func MultipleOfInt(factor int64) core.IntValidator {
	return func(value int64) core.Error {
		if value%factor != 0 {
			return &errors.IntParamError{
				BasicError: multipleOfError,
				ParamKey:   errors.ParamKeyFactor,
				ParamValue: factor,
			}
		}
		return nil
	}
}
