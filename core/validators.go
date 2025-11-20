package core

// IntValidator validates integer fields.
type IntValidator func(int64) Error

// FloatValidator validates float fields.
type FloatValidator func(float64) Error
