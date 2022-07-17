package rest

import (
	"context"
	"github.com/gorilla/mux"
	"loginMicroservice/app/internal/datasource"
	"loginMicroservice/app/internal/logger"
	handlers2 "loginMicroservice/app/internal/transport/rest/handlers"
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

func NewRestServer(
	ctx context.Context,
	log logger.Logger,
	url string,
	dbSourcer datasource.DbSourcer,
) *RestServer {
	router := mux.NewRouter()
	router.Handle("/register", handlers2.NewRegistrationHandler(ctx, dbSourcer, log)).Methods(http.MethodPost)
	router.Handle("/authorization", handlers2.NewAuthorizeHandler(ctx, dbSourcer, log)).Methods(http.MethodPost)

	server := &http.Server{
		Addr:         url,
		Handler:      router,
		WriteTimeout: time.Second,
		ReadTimeout:  time.Second,
	}

	return &RestServer{
		log:        log,
		ctx:        ctx,
		httpServer: server,
	}
}

func (rs *RestServer) Run(cancel func()) {
	if err := rs.httpServer.ListenAndServe(); err != nil {
		rs.log.WithField("server error", err.Error()).Error()
		cancel()
	}
}

func (rs *RestServer) Shutdown() {
	if err := rs.httpServer.Shutdown(rs.ctx); err != nil {
		rs.log.WithField("shutdown err", err)
	}
}
