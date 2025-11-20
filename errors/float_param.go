package errors

import "github.com/mailru/easyjson/jwriter"

// FloatParamError provides additional information about the validation constraint,
// using a float parameter.
type FloatParamError struct {
	BasicError
	ParamKey   string
	ParamValue float64
}

// Error returns the human-readable message for the failed validation rule.
func (e *FloatParamError) Error() string {
	return e.Message
}

// MarshalEasyJSON writes the structured JSON representation of the validation error.
func (e *FloatParamError) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('{')
	w.RawString(keyCode)
	w.String(string(e.Code))
	w.RawString(keyMessage)
	w.String(e.Message)
	w.RawString(e.ParamKey)
	w.Float64(e.ParamValue)
	w.RawByte('}')
}
