package constraints

import (
	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var requiredError = &errors.BasicError{
	Code:    core.RuleRequired,
	Message: "field is required",
}

var notNullError = &errors.BasicError{
	Code:    core.RuleNotNull,
	Message: "value must not be null",
}

func RequiredAny(absent bool) core.Error {
	if absent {
		return requiredError
	}
	return nil
}

func NotNullAny(null bool) core.Error {
	if null {
		return notNullError
	}
	return nil
}
