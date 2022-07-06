package server

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
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
	handlers map[string]http.Handler,
) *RestServer {
	r := mux.NewRouter()
	r.Use()
	for key, handler := range handlers {
		r.Handle(key, handler)
	}

	server := &http.Server{
		Addr:    url,
		Handler: r,
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
