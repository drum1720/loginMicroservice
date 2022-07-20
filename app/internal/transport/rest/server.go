package rest

import (
	"context"
	"loginMicroservice/app/internal/configs"
	"loginMicroservice/app/internal/logger"
	"net/http"
	"time"
)

type Server struct {
	ctx        context.Context
	httpServer *http.Server
	log        logger.Logger
	cfg        configs.Configure
	router     http.Handler
}

func NewRestServer(
	ctx context.Context,
	log logger.Logger,
	cfg configs.Configure,
	router http.Handler,
) *Server {
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
