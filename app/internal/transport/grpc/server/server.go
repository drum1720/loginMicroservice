package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
)

type Server interface {
	Run()
}

type RestServer struct {
	ctx        context.Context
	url        string
	httpServer *http.Server
	log        *logrus.Logger
}

func NewRestServer(ctx context.Context,
	log *logrus.Logger,
	url string,
	handler http.Handler,
) *RestServer {

	server := &http.Server{
		Addr:         url,
		Handler:      handler,
		WriteTimeout: time.Second,
		ReadTimeout:  time.Second,
	}

	return &RestServer{
		url:        url,
		log:        log,
		ctx:        ctx,
		httpServer: server,
	}
}

func (rs *RestServer) Run() {
	go log.Fatal(rs.httpServer.ListenAndServe())
}
