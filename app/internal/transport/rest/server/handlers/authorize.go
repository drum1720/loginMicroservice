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

type AuthorizeHandler struct {
	ctx context.Context
	db  *posgresql.DbConnectionPool
	log *logrus.Logger
}

func NewAuthorizeHandler(
	ctx context.Context,
	db *posgresql.DbConnectionPool,
	log *logrus.Logger,
) *AuthorizeHandler {
	return &AuthorizeHandler{
		ctx: ctx,
		db:  db,
		log: log,
	}
}

func (ah AuthorizeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user core.User

	if err := request.ParseData(r.Body, &user); err != nil {
		ah.log.WithField("err", err.Error()).Warning("parse data err")
		http.Error(w, err.Error(), err.GetStatusCode())
		return
	}

	if ok, err := ah.db.UserValid(ah.ctx, &user); err != nil || !ok {
		ah.log.WithField("error", err.Error()).Info("user valid err")
		http.Error(w, fmt.Sprintf("user or pass not valid"), err.GetStatusCode())
		return
	}

	response.NewAuthorizeResponse(user).Write(w)

}
