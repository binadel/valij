package results

import (
	"github.com/binadel/vali/core"
	"github.com/mailru/easyjson/jwriter"
)

// StringResult represents the validation result for a string field.
type StringResult struct {
	Path   core.FieldPath
	Value  string
	Errors []core.Error
}

// IsValid returns whether the field has passed validation test.
func (r StringResult) IsValid() bool {
	return len(r.Errors) == 0
}

// MarshalEasyJSON writes the json representation to the output.
func (r StringResult) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('{')

	w.RawString(keyPath)
	r.Path.MarshalEasyJSON(w)

	w.RawString(keyValue)
	w.String(r.Value)

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
