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

var minLengthError = errors.BasicError{
	Code:    core.RuleMinLength,
	Message: "value length must be at least {{minLength}}",
}

var maxLengthError = errors.BasicError{
	Code:    core.RuleMaxLength,
	Message: "value length must be at most {{maxLength}}",
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
