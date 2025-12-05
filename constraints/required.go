package constraints

import (
	"github.com/binadel/valij/core"
	"github.com/binadel/valij/errors"
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
