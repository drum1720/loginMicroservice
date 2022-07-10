package handlers

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"loginMicroservice/app/internal/core"
	"loginMicroservice/app/internal/datasource/posgresql"
	"loginMicroservice/app/internal/transport/rest/server/request"
	"loginMicroservice/app/internal/transport/rest/server/response"
	"net/http"
)

type RegistrationHandler struct {
	ctx context.Context
	db  *posgresql.DbConnectionPool
	log *logrus.Logger
}

func NewRegistrationHandler(
	ctx context.Context,
	db *posgresql.DbConnectionPool,
	log *logrus.Logger,
) *RegistrationHandler {
	return &RegistrationHandler{
		ctx: ctx,
		db:  db,
		log: log,
	}
}

func (rh RegistrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user core.User

	if err := request.ParseData(r.Body, &user); err != nil {
		rh.log.WithField("err", err.Error()).Warning("parse data err")
		http.Error(w, err.Error(), err.GetStatusCode())
		return
	}

	if ok, err := rh.db.UserExist(rh.ctx, user); err != nil || ok {
		rh.log.WithField("error", err.Error()).Info("user not created")
		http.Error(w, fmt.Sprintf("user not created: err: %s", err.Error()), err.GetStatusCode())
		return
	}

	if err := rh.db.InsertUser(rh.ctx, user); err != nil {
		rh.log.WithField("err", err).Warning()
		http.Error(w, err.Error(), err.GetStatusCode())
		return
	}

	response.NewRegistrationResponse(user).Write(w)
}
