package handlers

import (
	"context"
	"fmt"
	"loginMicroservice/app/internal/datasource"
	"loginMicroservice/app/internal/logger"
	"loginMicroservice/app/internal/transport/rest/common"
	"loginMicroservice/app/internal/transport/rest/handlers/request"
	"loginMicroservice/app/internal/transport/rest/handlers/response"
	"net/http"
)

type RegistrationHandler struct {
	ctx context.Context
	db  datasource.DbSourcer
	log logger.Logger
}

func NewRegistrationHandler(
	ctx context.Context,
	db datasource.DbSourcer,
	log logger.Logger,
) *RegistrationHandler {
	return &RegistrationHandler{
		ctx: ctx,
		db:  db,
		log: log,
	}
}

func (h RegistrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user request.Registration

	if err := common.ParseData(r.Body, &user); err != nil {
		h.log.WithField("err", err.Error()).Error("parse data err")
		http.Error(w, err.Error(), err.GetStatusCode())
		return
	}

	if ok, err := h.db.UserExist(h.ctx, user.User); err != nil || ok {
		h.log.WithField("error", err.Error()).Info("user not created")
		http.Error(w, fmt.Sprintf("user not created: err: %s", err.Error()), err.GetStatusCode())
		return
	}

	if err := h.db.InsertUser(h.ctx, user.User, user.Password); err != nil {
		h.log.WithField("err", err).Error()
		http.Error(w, err.Error(), err.GetStatusCode())
		return
	}

	response.NewRegistrationResponse(user.User).Write(w)
	h.log.WithFields(logger.Fields{"user": user.User}).Info("register to service")
}
