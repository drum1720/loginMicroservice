package handlers

import (
	"context"
	"github.com/sirupsen/logrus"
	"loginMicroservice/app/internal/datasource/posgresql"
	"net/http"
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
	//TODO implement me
	panic("implement me")
}
