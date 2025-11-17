package errors

import (
	"github.com/binadel/vali/core"
	"github.com/mailru/easyjson/jwriter"
)

const (
	keyCode    = "\"code\":"
	keyMessage = ",\"message\":"
)

type BasicError struct {
	Code    core.Rule
	Message string
}

func (e *BasicError) Error() string {
	return e.Message
}

func (e *BasicError) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('{')
	w.RawString(keyCode)
	w.String(string(e.Code))
	w.RawString(keyMessage)
	w.String(e.Message)
	w.RawByte('}')
}
