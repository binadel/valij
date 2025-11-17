package vali

import (
	"github.com/binadel/vali/core"
	"github.com/binadel/vali/validations"
)

func Int(path ...string) validations.IntValidation {
	return validations.IntValidation{
		Path: core.Field(path),
	}
}
