package validations

import (
	"github.com/binadel/vali/constraints"
	"github.com/binadel/vali/validations/results"
	"github.com/google/uuid"
)

type UuidValidation struct {
	stringValidation StringValidation
	version          uuid.Version
}

func (v UuidValidation) Version(version byte) UuidValidation {
	if version >= 1 && version <= 8 {
		v.version = uuid.Version(version)
	}
	return v
}

// Validate applies the validations constraints to the field value and returns the result.
func (v UuidValidation) Validate(value string) results.FormatResult[uuid.UUID] {
	var result results.FormatResult[uuid.UUID]
	result.StringResult = v.stringValidation.Validate(value)

	if len(result.Errors) == 0 {
		if id, err := constraints.ParseUuid(value, v.version); err != nil {
			result.Errors = append(result.Errors, err)
		} else {
			result.Format = id
		}
	}

	return result
}
