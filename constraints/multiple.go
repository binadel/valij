package constraints

import (
	"github.com/binadel/valij/core"
	"github.com/binadel/valij/errors"
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
