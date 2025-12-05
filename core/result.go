package core

import (
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jwriter"
)

const (
	keyPath   = `"path":`
	keyValue  = `,"value":`
	keyErrors = `,"errors":`
)

type value interface {
	easyjson.Optional
	easyjson.Marshaler
}

// Result represents the validation result for the field.
type Result[T value] struct {
	Path   FieldPath
	Value  T
	Errors []Error
}

// MarshalEasyJSON implements easyjson.Marshaler.
func (r Result[T]) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('{')

	w.RawString(keyPath)
	r.Path.MarshalEasyJSON(w)

	if r.Value.IsDefined() {
		w.RawString(keyValue)
		r.Value.MarshalEasyJSON(w)
	}

	w.RawString(keyErrors)
	w.RawByte('[')
	for i, err := range r.Errors {
		if i > 0 {
			w.RawByte(',')
		}
		err.MarshalEasyJSON(w)
	}
	w.RawByte(']')

	w.RawByte('}')
}
