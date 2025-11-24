package errors

import "github.com/mailru/easyjson/jwriter"

const (
	ParamKeyMinimum   = ",\"minimum\":"
	ParamKeyMaximum   = ",\"maximum\":"
	ParamKeyFactor    = ",\"factor\":"
	ParamKeyLength    = ",\"length\":"
	ParamKeyMinLength = ",\"minLength\":"
	ParamKeyMaxLength = ",\"maxLength\":"
	ParamKeyVersion   = ",\"version\":"
)

// IntParamError provides additional information about the validation constraint,
// using an integer parameter.
type IntParamError struct {
	BasicError
	ParamKey   string
	ParamValue int64
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

// FloatParamError provides additional information about the validation constraint,
// using a float parameter.
type FloatParamError struct {
	BasicError
	ParamKey   string
	ParamValue float64
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
