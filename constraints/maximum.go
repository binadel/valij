package constraints

import (
	"github.com/binadel/valij/core"
	"github.com/binadel/valij/errors"
)

var maximumError = errors.BasicError{
	Code:    core.RuleMaximum,
	Message: "value must be less than or equal to {{maximum}}",
}

var exclusiveMaximumError = errors.BasicError{
	Code:    core.RuleExclusiveMaximum,
	Message: "value must be less than {{maximum}}",
}

// MaximumInt applies maximum constraint to an integer.
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

// MaximumFloat applies maximum constraint to a float.
func MaximumFloat(maximum float64) core.FloatValidator {
	return func(value float64) core.Error {
		if value > maximum {
			return &errors.FloatParamError{
				BasicError: maximumError,
				ParamKey:   errors.ParamKeyMaximum,
				ParamValue: maximum,
			}
		}
		return nil
	}
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

// ExclusiveMaximumFloat applies exclusive maximum constraint to a float.
func ExclusiveMaximumFloat(maximum float64) core.FloatValidator {
	return func(value float64) core.Error {
		if value >= maximum {
			return &errors.FloatParamError{
				BasicError: exclusiveMaximumError,
				ParamKey:   errors.ParamKeyMaximum,
				ParamValue: maximum,
			}
		}
		return nil
	}
}
