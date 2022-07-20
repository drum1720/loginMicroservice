package common

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	error2 "loginMicroservice/app/internal/transport/rest/error"
	"loginMicroservice/app/internal/transport/rest/handlers/request"
	"net/http"
)

func ParseData(body io.ReadCloser, essence request.Validator) *error2.Error {
	buffer, err := io.ReadAll(body)
	if err != nil {
		return error2.NewError(err, http.StatusInternalServerError)
	}

	if err = json.Unmarshal(buffer, &essence); err != nil {
		e, ok := err.(*json.UnmarshalTypeError)
		if ok {
			return error2.NewError(
				errors.Errorf("%s: empty type - '%s'. Need type string", e.Field, e.Value),
				http.StatusBadRequest,
			)
		}
		return error2.NewError(errors.New("unmarshal err"), http.StatusInternalServerError)
	}

	if err = essence.Validate(); err != nil {
		return error2.NewError(err, http.StatusBadRequest)
	}

	return nil
}
