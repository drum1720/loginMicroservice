package request

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"loginMicroservice/app/internal/core"
	"loginMicroservice/app/internal/transport/rest/server/response"
	"net/http"
)

func ParseData(body io.ReadCloser, essence core.Validater) *response.Error {
	buffer, err := io.ReadAll(body)
	if err != nil {
		return response.NewError(err, http.StatusInternalServerError)
	}

	if err = json.Unmarshal(buffer, &essence); err != nil {
		e, ok := err.(*json.UnmarshalTypeError)
		if ok {
			return response.NewError(errors.Errorf("%s: empty type - '%s'. Need type string", e.Field, e.Value), http.StatusBadRequest)
		}
	}
	if err = essence.Validate(); err != nil {
		return response.NewError(err, http.StatusBadRequest)
	}

	return nil
}
