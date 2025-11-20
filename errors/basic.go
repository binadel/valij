package errors

import (
	"github.com/binadel/vali/core"
	"github.com/mailru/easyjson/jwriter"
)

const (
	keyCode    = "\"code\":"
	keyMessage = ",\"message\":"
)

// BasicError provides error code and message.
// Code identifies the validation rule that failed.
// Message is the human-readable description of the error.
type BasicError struct {
	Code    core.Rule
	Message string
}

// Error returns the human-readable message for the failed validation rule.
func (e *BasicError) Error() string {
	return e.Message
}

// MarshalEasyJSON writes the structured JSON representation of the validation error.
func (e *BasicError) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('{')
	w.RawString(keyCode)
	w.String(string(e.Code))
	w.RawString(keyMessage)
	w.String(e.Message)
	w.RawByte('}')
}
