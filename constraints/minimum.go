package constraints

import (
	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var minimumError = errors.BasicError{
	Code:    core.RuleMinimum,
	Message: "value must be greater than or equal to {{minimum}}",
}

var exclusiveMinimumError = errors.BasicError{
	Code:    core.RuleExclusiveMinimum,
	Message: "value must be greater than {{minimum}}",
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

// MinimumFloat applies minimum constraint to a float.
func MinimumFloat(minimum float64) core.FloatValidator {
	return func(value float64) core.Error {
		if value < minimum {
			return &errors.FloatParamError{
				BasicError: minimumError,
				ParamKey:   errors.ParamKeyMinimum,
				ParamValue: minimum,
			}
		}
		return nil
	}
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

// ExclusiveMinimumFloat applies exclusive minimum constraint to a float.
func ExclusiveMinimumFloat(minimum float64) core.FloatValidator {
	return func(value float64) core.Error {
		if value <= minimum {
			return &errors.FloatParamError{
				BasicError: exclusiveMinimumError,
				ParamKey:   errors.ParamKeyMinimum,
				ParamValue: minimum,
			}
		}
		return nil
	}
}
