package response

import (
	"encoding/json"
)

type ResponseErr struct {
	err        error
	statusCode int
}

func NewResponseErr(err error, status int) *ResponseErr {
	return &ResponseErr{
		err:        err,
		statusCode: status}
}

func (e ResponseErr) Bytes() []byte {
	result, _ := json.Marshal(e)
	return result
}

func (e ResponseErr) Error() string {
	return e.err.Error()
}

func (e ResponseErr) GetStatusCode() int {
	return e.statusCode
}
