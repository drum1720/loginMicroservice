package response

import (
	"encoding/json"
	"loginMicroservice/app/internal/core"
	"net/http"
	"time"
)

type AuthorizeResponse struct {
	User      core.User
	Authorize bool
}

func NewAuthorizeResponse(user core.User) *AuthorizeResponse {
	return &AuthorizeResponse{
		User:      user,
		Authorize: false,
	}
}

func (ar AuthorizeResponse) Write(w http.ResponseWriter) {
	ar.User.LastVisit = time.Now()
	ar.User.Password = "*****"
	ar.Authorize = true

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(ar)

	w.Write(data)
}
