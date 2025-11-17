package constraints

import (
	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var minError = errors.BasicError{
	Code:    core.RuleMin,
	Message: "field value is less than allowed minimum",
}

func MinInt(min int64) core.IntValidator {
	return func(value int64) core.Error {
		if value < min {
			return &errors.IntParamError{
				BasicError: minError,
				ParamKey:   errors.ParamKeyLimit,
				ParamValue: min,
			}
		}
		return nil
	}
}
