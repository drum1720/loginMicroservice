package handlers

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"loginMicroservice/app/internal/core"
	"loginMicroservice/app/internal/datasource/posgresql"
	"net/http"
	"time"
)

type Authorization struct {
	ctx context.Context
	db  *posgresql.DbConnectionPool
	log *logrus.Logger
}

func NewAuthorization(
	ctx context.Context,
	db *posgresql.DbConnectionPool,
	log *logrus.Logger,
) *Authorization {
	return &Authorization{
		ctx: ctx,
		db:  db,
		log: log,
	}
}

func (a Authorization) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user core.User
	buffer, err := io.ReadAll(r.Body)
	if err != nil {
		a.log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.Unmarshal(buffer, &user)
	if err = user.Validation(); err != nil {
		a.log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = a.db.DbPool.Exec(a.ctx, "INSERT INTO \"auf\" (\"login\", \"pass\", \"last_visit\") VALUES ($1,$2,$3)", user.User, user.Password, time.Now())
	if err != nil {
		a.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(user)
	if err != nil {
		a.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
