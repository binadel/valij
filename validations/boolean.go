package validations

import (
	"github.com/binadel/jsontypes"
	"github.com/binadel/valij/constraints"
	"github.com/binadel/valij/core"
)

// BooleanValidation defines the validation constraints for a JSON boolean field.
type BooleanValidation struct {
	Path     core.FieldPath
	required bool
	notNull  bool
}

func (v BooleanValidation) Required() BooleanValidation {
	v.required = true
	return v
}

func (v BooleanValidation) NotNull() BooleanValidation {
	v.notNull = true
	return v
}

func (v BooleanValidation) Validate(value jsontypes.Boolean) core.Result[jsontypes.Boolean] {
	var errors []core.Error
	if v.required {
		if err := constraints.RequiredAny(!value.Present); err != nil {
			errors = append(errors, err)
		}
	}
	if v.notNull {
		if err := constraints.NotNullAny(!value.Valid); err != nil {
			errors = append(errors, err)
		}
	}

	return core.Result[jsontypes.Boolean]{
		Path:   v.Path,
		Value:  value,
		Errors: errors,
	}
}
