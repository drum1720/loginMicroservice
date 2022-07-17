package handlers

import (
	"context"
	"fmt"
	"loginMicroservice/app/internal/core"
	"loginMicroservice/app/internal/datasource"
	"loginMicroservice/app/internal/logger"
	"loginMicroservice/app/internal/transport/rest/request"
	"loginMicroservice/app/internal/transport/rest/response"
	"net/http"
)

type AuthorizeHandler struct {
	ctx context.Context
	db  datasource.DbSourcer
	log logger.Logger
}

func NewAuthorizeHandler(
	ctx context.Context,
	db datasource.DbSourcer,
	log logger.Logger,
) *AuthorizeHandler {
	return &AuthorizeHandler{
		ctx: ctx,
		db:  db,
		log: log,
	}
}

func (h AuthorizeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user core.User

	if err := request.ParseData(r.Body, &user); err != nil {
		h.log.WithField("err", err.Error()).Error("parse data err")
		http.Error(w, err.Error(), err.GetStatusCode())
		return
	}

	if ok, err := h.db.UserValid(h.ctx, &user); err != nil || !ok {
		h.log.WithField("error", err.Error()).Info("user valid err")
		http.Error(w, fmt.Sprintf("user or pass not valid"), err.GetStatusCode())
		return
	}

	response.NewAuthorizeResponse(user).Write(w)
	h.log.WithFields(logger.Fields{"user": user.User}).Info("visit to service")

}
