package errors

import "github.com/mailru/easyjson/jwriter"

// IntParamError provides additional information about the validation constraint,
// using an integer parameter.
type IntParamError struct {
	BasicError
	ParamKey   string
	ParamValue int64
}

// Error returns the human-readable message for the failed validation rule.
func (e *IntParamError) Error() string {
	return e.Message
}

// MarshalEasyJSON writes the structured JSON representation of the validation error.
func (e *IntParamError) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('{')
	w.RawString(keyCode)
	w.String(string(e.Code))
	w.RawString(keyMessage)
	w.String(e.Message)
	w.RawString(e.ParamKey)
	w.Int64(e.ParamValue)
	w.RawByte('}')
}
