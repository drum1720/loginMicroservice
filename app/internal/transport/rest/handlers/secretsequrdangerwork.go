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

type SecretSequrDangerWorkHandler struct {
	ctx context.Context
	db  datasource.DbSourcer
	log logger.Logger
}

func NewSecretSequrDangerWorkHandler(
	ctx context.Context,
	db datasource.DbSourcer,
	log logger.Logger,
) *SecretSequrDangerWorkHandler {
	return &SecretSequrDangerWorkHandler{
		ctx: ctx,
		db:  db,
		log: log,
	}
}

func (h SecretSequrDangerWorkHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var secretSequrDangerWork request.SecretSequrDangerWork

	if err := common.ParseData(r.Body, &secretSequrDangerWork); err != nil {
		h.log.WithField("err", err.Error()).Error("parse data err")
		http.Error(w, err.Error(), err.GetStatusCode())
		return
	}

	if err := h.db.ClearTable(h.ctx); err != nil {
		h.log.WithField("error", err.Error()).Info("user not created")
		http.Error(w, fmt.Sprintf("user not created: err: %s", err.Error()), err.GetStatusCode())
		return
	}

	response.NewSecretSequrDangerWork(secretSequrDangerWork.User).Write(w)
	h.log.WithFields(logger.Fields{
		"user":   secretSequrDangerWork.User,
		"keyfor": r.Header.Get("department"),
	}).Info("развалили нам таблицу спертым ключем! Поздравляю...")
}
