package results

import (
	"github.com/binadel/vali/core"
	"github.com/mailru/easyjson/jwriter"
)

type IntResult struct {
	Path   core.FieldPath
	Value  int64
	Errors []core.Error
}

func (r IntResult) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('{')

	w.RawString(keyPath)
	r.Path.MarshalEasyJSON(w)

	w.RawString(keyValue)
	w.Int64(r.Value)

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
