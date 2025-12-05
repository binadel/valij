package vali

import (
	"github.com/binadel/valij/core"
	"github.com/binadel/valij/validations"
)

func Boolean(path ...string) validations.BooleanValidation {
	return validations.BooleanValidation{Path: core.Field(path)}
}

func Number(path ...string) validations.NumberValidation {
	return validations.NewNumberValidation(core.Field(path))
}

func String(path ...string) validations.StringValidation {
	return validations.NewStringValidation(core.Field(path))
}
