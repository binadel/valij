package validations

import (
	"github.com/binadel/vali/constraints"
	"github.com/binadel/vali/core"
	"github.com/binadel/vali/validations/results"
)

type IntValidation struct {
	Path        core.FieldPath
	constraints []core.IntValidator
}

func (v IntValidation) Minimum(minimum int64) IntValidation {
	v.constraints = append(v.constraints, constraints.MinimumInt(minimum))
	return v
}

func (v IntValidation) ExclusiveMinimum(exclusiveMinimum int64) IntValidation {
	v.constraints = append(v.constraints, constraints.ExclusiveMinimumInt(exclusiveMinimum))
	return v
}

func (v IntValidation) Maximum(maximum int64) IntValidation {
	v.constraints = append(v.constraints, constraints.MaximumInt(maximum))
	return v
}

func (v IntValidation) ExclusiveMaximum(exclusiveMaximum int64) IntValidation {
	v.constraints = append(v.constraints, constraints.ExclusiveMaximumInt(exclusiveMaximum))
	return v
}

func (v IntValidation) Validate(value int64) results.IntResult {
	var errors []core.Error
	for _, constraint := range v.constraints {
		if err := constraint(value); err != nil {
			errors = append(errors, err)
		}
	}
	return results.IntResult{
		Path:   v.Path,
		Value:  value,
		Errors: errors,
	}
}
