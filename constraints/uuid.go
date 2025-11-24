package constraints

import (
	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
	"github.com/google/uuid"
)

var uuidError = &errors.BasicError{
	Code:    core.RuleUuid,
	Message: "value must be a valid UUID",
}

// ParseUuid parses the given string as a UUID.
func ParseUuid(value string) (uuid.UUID, core.Error) {
	if id, err := uuid.Parse(value); err == nil {
		return id, nil
	}
	return uuid.Nil, uuidError
}
