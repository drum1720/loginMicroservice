package rest

import (
	"context"
	"github.com/gorilla/mux"
	"loginMicroservice/app/internal/configs"
	"loginMicroservice/app/internal/datasource"
	"loginMicroservice/app/internal/logger"
	handlers2 "loginMicroservice/app/internal/transport/rest/handlers"
	"net/http"
	"time"
)

type Server struct {
	ctx        context.Context
	httpServer *http.Server
	log        logger.Logger
	cfg        configs.Configure
}

func NewRestServer(
	ctx context.Context,
	log logger.Logger,
	cfg configs.Configure,
	dbSourcer datasource.DbSourcer,
) *Server {
	router := mux.NewRouter()
	router.Handle("/register", handlers2.NewRegistrationHandler(ctx, dbSourcer, log)).Methods(http.MethodPost)
	router.Handle("/authorization", handlers2.NewAuthorizeHandler(ctx, dbSourcer, log, cfg)).Methods(http.MethodPost)

	server := &http.Server{
		Addr:         cfg.GetUrl(),
		Handler:      router,
		WriteTimeout: time.Second,
		ReadTimeout:  time.Second,
	}

	return &Server{
		log:        log,
		ctx:        ctx,
		httpServer: server,
		cfg:        cfg,
	}
}

func (rs *Server) Run(cancel func()) {
	if err := rs.httpServer.ListenAndServe(); err != nil {
		rs.log.WithField("server error", err.Error()).Error()
		cancel()
	}
}

func (rs *Server) Shutdown() {
	if err := rs.httpServer.Shutdown(rs.ctx); err != nil {
		rs.log.WithField("shutdown err", err)
	}
}
