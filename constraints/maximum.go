package constraints

import (
	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var maximumError = errors.BasicError{
	Code:    core.RuleMaximum,
	Message: "value must be less than or equal to {{maximum}}",
}

func MaximumInt(maximum int64) core.IntValidator {
	return func(value int64) core.Error {
		if value > maximum {
			return &errors.IntParamError{
				BasicError: maximumError,
				ParamKey:   errors.ParamKeyMaximum,
				ParamValue: maximum,
			}
		}
		return nil
	}
}
