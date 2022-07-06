package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type RestServer struct {
	ctx        context.Context
	url        string
	httpServer *http.Server
	log        *logrus.Logger
	handler    http.Handler
}

func NewRestServer(ctx context.Context, log *logrus.Logger, url string, handler http.Handler) *RestServer {
	srv := &http.Server{
		Handler: handler,
		Addr:    url,
	}

	return &RestServer{
		url:        url,
		log:        log,
		ctx:        ctx,
		handler:    handler,
		httpServer: srv,
	}
}

func (rs *RestServer) Run() {
	go log.Fatal(rs.httpServer.ListenAndServe())
}
