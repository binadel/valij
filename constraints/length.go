package constraints

import (
	"unicode/utf8"

	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var lengthError = errors.BasicError{
	Code:    core.RuleLength,
	Message: "value length must be equal to {{length}}",
}

// LengthString applies length constraint to a string.
func LengthString(expectedLength int) core.StringValidator {
	return func(value string) core.Error {
		if length := utf8.RuneCountInString(value); length != expectedLength {
			return &errors.IntParamError{
				BasicError: lengthError,
				ParamKey:   errors.ParamKeyLength,
				ParamValue: int64(expectedLength),
			}
		}
		return nil
	}
}
