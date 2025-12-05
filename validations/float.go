package validations

import (
	"github.com/binadel/valij/constraints"
	"github.com/binadel/valij/core"
	"github.com/binadel/valij/validations/results"
)

// FloatValidation defines the validation constraints for an integer field.
type FloatValidation struct {
	Path        core.FieldPath
	constraints []core.FloatValidator
}

// Minimum adds a minimum constraint to the field.
func (v FloatValidation) Minimum(minimum float64) FloatValidation {
	v.constraints = append(v.constraints, constraints.MinimumFloat(minimum))
	return v
}

// Maximum adds a maximum constraint to the field.
func (v FloatValidation) Maximum(maximum float64) FloatValidation {
	v.constraints = append(v.constraints, constraints.MaximumFloat(maximum))
	return v
}

// ExclusiveMinimum adds an exclusive minimum constraint to the field.
func (v FloatValidation) ExclusiveMinimum(exclusiveMinimum float64) FloatValidation {
	v.constraints = append(v.constraints, constraints.ExclusiveMinimumFloat(exclusiveMinimum))
	return v
}

// ExclusiveMaximum adds an exclusive maximum constraint to the field.
func (v FloatValidation) ExclusiveMaximum(exclusiveMaximum float64) FloatValidation {
	v.constraints = append(v.constraints, constraints.ExclusiveMaximumFloat(exclusiveMaximum))
	return v
}

// Validate applies the validations constraints to the field value and returns the result.
func (v FloatValidation) Validate(value float64) results.FloatResult {
	var errors []core.Error
	for _, constraint := range v.constraints {
		if err := constraint(value); err != nil {
			errors = append(errors, err)
		}
	}
	return results.FloatResult{
		Path:   v.Path,
		Value:  value,
		Errors: errors,
	}
}
