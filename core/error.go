package core

import "github.com/mailru/easyjson"

type Error interface {
	error
	easyjson.Marshaler
}
