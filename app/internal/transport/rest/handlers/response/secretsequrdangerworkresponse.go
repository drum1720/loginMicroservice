package response

import (
	"encoding/json"
	"net/http"
	"time"
)

type SecretSequrDangerWork struct {
	StatusCode int       `json:"-"`
	User       string    `json:"user"`
	DataTime   time.Time `json:"time"`
	Message    string    `json:"message"`
}

// NewSecretSequrDangerWork ...
func NewSecretSequrDangerWork(user string) *SecretSequrDangerWork {
	return &SecretSequrDangerWork{
		StatusCode: http.StatusOK,
		User:       user,
		Message:    "таблица auf очищена",
	}
}

// Write ...
func (rr SecretSequrDangerWork) Write(w http.ResponseWriter) {
	rr.DataTime = time.Now()

	w.WriteHeader(rr.StatusCode)
	data, _ := json.Marshal(&rr)
	w.Write(data)
}
