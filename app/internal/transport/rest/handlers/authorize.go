package handlers

import (
	"context"
	"fmt"
	"loginMicroservice/app/internal/configs"
	"loginMicroservice/app/internal/datasource"
	"loginMicroservice/app/internal/logger"
	"loginMicroservice/app/internal/transport/rest/common"
	"loginMicroservice/app/internal/transport/rest/handlers/request"
	"loginMicroservice/app/internal/transport/rest/handlers/response"
	"net/http"
)

type AuthorizeHandler struct {
	ctx context.Context
	db  datasource.DbSourcer
	log logger.Logger
	cfg configs.Configure
}

func NewAuthorizeHandler(
	ctx context.Context,
	db datasource.DbSourcer,
	log logger.Logger,
	cfg configs.Configure,
) *AuthorizeHandler {
	return &AuthorizeHandler{
		ctx: ctx,
		db:  db,
		log: log,
		cfg: cfg,
	}
}

func (h AuthorizeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user request.Authorize

	if err := common.ParseData(r.Body, &user); err != nil {
		h.log.WithField("err", err.Error()).Error("parse data err")
		http.Error(w, err.Error(), err.GetStatusCode())
		return
	}

	if ok, err := h.db.UserValid(h.ctx, user.User, user.Password); err != nil || !ok {
		h.log.WithField("error", err.Error()).Info("user valid err")
		http.Error(w, fmt.Sprintf("user or pass not valid"), err.GetStatusCode())
		return
	}

	response.NewAuthorizeResponse(user.User, h.cfg.GetJWT()).Write(w)
	h.log.WithFields(logger.Fields{"user": user.User}).Info("visit to service")
}
