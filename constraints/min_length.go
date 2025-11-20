package constraints

import (
	"unicode/utf8"

	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var minLengthError = errors.BasicError{
	Code:    core.RuleMinLength,
	Message: "value length must be at least {{minLength}}",
}

// MinLengthString applies min length constraint to a string.
func MinLengthString(minLength int) core.StringValidator {
	return func(value string) core.Error {
		if length := utf8.RuneCountInString(value); length < minLength {
			return &errors.IntParamError{
				BasicError: minLengthError,
				ParamKey:   errors.ParamKeyMinLength,
				ParamValue: int64(minLength),
			}
		}
		return nil
	}
}
