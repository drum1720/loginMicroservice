package response

import (
	"encoding/json"
	"net/http"
	"time"
)

type RegistrationResponse struct {
	StatusCode int       `json:"-"`
	User       string    `json:"user"`
	DataTime   time.Time `json:"time"`
}

func NewRegistrationResponse(user string) *RegistrationResponse {
	return &RegistrationResponse{
		StatusCode: http.StatusCreated,
		User:       user,
	}
}

func (rr RegistrationResponse) Write(w http.ResponseWriter) {
	rr.DataTime = time.Now()

	w.WriteHeader(rr.StatusCode)
	data, _ := json.Marshal(&rr) //косяк с маршалом!!!!
	w.Write(data)
}
