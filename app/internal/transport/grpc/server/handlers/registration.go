package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"loginMicroservice/app/internal/core"
	"loginMicroservice/app/internal/datasource/posgresql"
	"loginMicroservice/app/internal/transport/grpc/server/response"
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

	if err := parseData(r.Body, &user); err != nil {
		rh.log.Error(err.Error())
		http.Error(w, err.Error(), err.GetStatusCode())
		return
	}

	if rh.db.ExistUser(rh.ctx, user) {
		rh.log.Debugf("user not created: user %s already exists", user.User)
		http.Error(w, fmt.Sprintf("user not created: user %s already exists", user.User), http.StatusBadRequest)
		return
	}

	if err := rh.db.InsertUser(rh.ctx, user); err != nil {
		rh.log.Error(err)
		http.Error(w, err.Error(), err.GetStatusCode())
		return
	}

	response.NewRegistrationResponse(user).Write(w)
}

func parseData(body io.ReadCloser, essence core.Validater) *response.ResponseErr {
	buffer, err := io.ReadAll(body)
	if err != nil {
		return response.NewResponseErr(err, http.StatusInternalServerError)
	}

	json.Unmarshal(buffer, &essence)
	if err = essence.Validate(); err != nil {
		return response.NewResponseErr(err, http.StatusBadRequest)
	}

	return nil
}
