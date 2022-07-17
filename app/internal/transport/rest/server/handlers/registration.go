package handlers

import (
	"context"
	"fmt"
	"loginMicroservice/app/internal/core"
	"loginMicroservice/app/internal/datasource/posgresql"
	"loginMicroservice/app/internal/logger"
	"loginMicroservice/app/internal/transport/rest/server/request"
	"loginMicroservice/app/internal/transport/rest/server/response"
	"net/http"
)

type RegistrationHandler struct {
	ctx context.Context
	db  *posgresql.DbConnectionPool
	log logger.Logger
}

func NewRegistrationHandler(
	ctx context.Context,
	db *posgresql.DbConnectionPool,
	log logger.Logger,
) *RegistrationHandler {
	return &RegistrationHandler{
		ctx: ctx,
		db:  db,
		log: log,
	}
}

func (h RegistrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user core.User

	if err := request.ParseData(r.Body, &user); err != nil {
		h.log.WithField("err", err.Error()).Error("parse data err")
		http.Error(w, err.Error(), err.GetStatusCode())
		return
	}

	if ok, err := h.db.UserExist(h.ctx, user); err != nil || ok {
		h.log.WithField("error", err.Error()).Info("user not created")
		http.Error(w, fmt.Sprintf("user not created: err: %s", err.Error()), err.GetStatusCode())
		return
	}

	if err := h.db.InsertUser(h.ctx, user); err != nil {
		h.log.WithField("err", err).Error()
		http.Error(w, err.Error(), err.GetStatusCode())
		return
	}

	response.NewRegistrationResponse(user).Write(w)
	h.log.WithFields(logger.Fields{"user": user.User}).Info("register to service")
}
