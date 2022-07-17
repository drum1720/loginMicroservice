package error

import (
	"encoding/json"
)

type Error struct {
	err        error
	statusCode int
}

func NewError(err error, status int) *Error {
	return &Error{
		err:        err,
		statusCode: status}
}

func (e Error) Bytes() []byte {
	result, _ := json.Marshal(e)
	return result
}

func (e Error) Error() string {
	return e.err.Error()
}

func (e Error) GetStatusCode() int {
	return e.statusCode
}
