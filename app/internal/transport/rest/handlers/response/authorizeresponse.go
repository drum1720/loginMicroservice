package response

import (
	"encoding/json"
	"net/http"
)

type AuthorizeResponse struct {
	User     string
	TokenJWT string
}

func NewAuthorizeResponse(user string, token string) *AuthorizeResponse {
	return &AuthorizeResponse{
		User:     user,
		TokenJWT: token,
	}
}

func (ar AuthorizeResponse) Write(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(ar)

	w.Write(data)
}
