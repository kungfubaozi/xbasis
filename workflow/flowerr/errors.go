package flowerr

import "encoding/json"

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func err(code int64, message string) *Error {
	return &Error{Code: code, Message: message}
}

func FromError(err error) *Error {
	return &Error{Message: err.Error()}
}
