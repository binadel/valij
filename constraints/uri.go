package constraints

import (
	"net/url"

	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var uriError = &errors.BasicError{
	Code:    core.RuleUri,
	Message: "value must be a valid URI",
}

// ParseUuid parses the given string as a UUID.
func ParseUri(value string) (*url.URL, core.Error) {
	if uri, err := url.Parse(value); err == nil {
		return uri, nil
	}
	return nil, uriError
}
