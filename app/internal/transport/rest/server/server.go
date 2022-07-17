package server

import (
	"context"
	"loginMicroservice/app/internal/logger"
	"net/http"
	"time"
)

type Server interface {
	Run()
}

type RestServer struct {
	ctx        context.Context
	httpServer *http.Server
	log        logger.Logger
}

func NewRestServer(ctx context.Context,
	log logger.Logger,
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
		log:        log,
		ctx:        ctx,
		httpServer: server,
	}
}

func (rs *RestServer) Run() {
	go rs.log.Error(rs.httpServer.ListenAndServe())
}
