package constraints

import (
	"net/url"

	"github.com/binadel/valij/core"
	"github.com/binadel/valij/errors"
	"github.com/google/uuid"
)

var uuidError = &errors.BasicError{
	Code:    core.RuleUuid,
	Message: "value must be a valid UUID",
}

var uuidVersionError = errors.BasicError{
	Code:    core.RuleUuidVersion,
	Message: "value must be a valid UUID version {{version}}",
}

var uriError = &errors.BasicError{
	Code:    core.RuleUri,
	Message: "value must be a valid URI",
}

// ParseUuid parses the given string as a UUID.
func ParseUuid(value string, version uuid.Version) (uuid.UUID, core.Error) {
	if id, err := uuid.Parse(value); err == nil {
		if version == 0 || id.Version() == version {
			return id, nil
		}
		return uuid.Nil, &errors.IntParamError{
			BasicError: uuidVersionError,
			ParamKey:   errors.ParamKeyVersion,
			ParamValue: int64(version),
		}
	}
	return uuid.Nil, uuidError
}

// ParseUri parses the given string as a URI.
func ParseUri(value string) (*url.URL, core.Error) {
	if uri, err := url.Parse(value); err == nil {
		return uri, nil
	}
	return nil, uriError
}
