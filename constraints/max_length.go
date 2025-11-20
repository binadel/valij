package constraints

import (
	"unicode/utf8"

	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var maxLengthError = errors.BasicError{
	Code:    core.RuleMaxLength,
	Message: "value length must be at most {{maxLength}}",
}

// MaxLengthString applies max length constraint to a string.
func MaxLengthString(maxLength int) core.StringValidator {
	return func(value string) core.Error {
		if length := utf8.RuneCountInString(value); length > maxLength {
			return &errors.IntParamError{
				BasicError: maxLengthError,
				ParamKey:   errors.ParamKeyMaxLength,
				ParamValue: int64(maxLength),
			}
		}
		return nil
	}
}
