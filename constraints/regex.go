package constraints

import (
	"regexp"

	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var regexError = &errors.BasicError{
	Code:    core.RuleRegex,
	Message: "value must a valid regular expression",
}

// ParseRegex parses the given string as a regular expression.
func ParseRegex(value string) (*regexp.Regexp, core.Error) {
	regex, err := regexp.Compile(value)
	if err != nil {
		return nil, regexError
	}
	return regex, nil
}
