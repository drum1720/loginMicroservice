package response

import (
	"encoding/json"
	"loginMicroservice/app/internal/core"
	"net/http"
	"time"
)

type registrationResponse struct {
	statusCode int
	user       core.User
	dataTime   time.Time
}

func NewRegistrationResponse(user core.User) *registrationResponse {
	return &registrationResponse{
		statusCode: http.StatusCreated,
		user:       user,
	}
}

func (rr registrationResponse) Write(w http.ResponseWriter) {
	rr.user.LastVisit = time.Now()
	rr.user.Password = "*****"

	w.WriteHeader(rr.statusCode)
	data, _ := json.Marshal(rr.user)
	w.Write(data)
}
