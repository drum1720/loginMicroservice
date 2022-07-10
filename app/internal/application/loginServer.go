package application

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"loginMicroservice/app/internal/core"
	"loginMicroservice/app/internal/datasource/posgresql"
	"loginMicroservice/app/internal/log"
	"loginMicroservice/app/internal/transport/rest/server"
	"loginMicroservice/app/internal/transport/rest/server/handlers"
	"net/http"
	"os"
)

type loginServer struct {
	log              *logrus.Logger
	ctx              context.Context
	cfg              *core.Cfg
	dbConnectionPool *posgresql.DbConnectionPool
	restServer       *server.RestServer
}

// NewLoginServer ...
func NewLoginServer() *loginServer {
	return &loginServer{}
}

// Init ...
func (ls *loginServer) Init() {
	ls.log = log.NewLogger()
	ls.ctx = context.Background()

	var err error
	ls.cfg, err = core.InitCfg()
	if err != nil {
		ls.log.WithField("config err", err).Error()
		os.Exit(1)
	}

	ls.dbConnectionPool, err = posgresql.NewDbConnectionPool(ls.cfg.GetDsnPG())
	if err != nil {
		ls.log.WithField("can't connect to database err", err.Error()).Error()
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.Handle("/register", handlers.NewRegistrationHandler(ls.ctx, ls.dbConnectionPool, ls.log)).Methods(http.MethodPost)
	router.Handle("/authorization", handlers.NewAuthorizeHandler(ls.ctx, ls.dbConnectionPool, ls.log)).Methods(http.MethodPost)

	ls.restServer = server.NewRestServer(ls.ctx, ls.log, ls.cfg.GetUrl(), router)
}

// Run ...
func (ls *loginServer) Run() {
	if err := ls.dbConnectionPool.Ping(ls.ctx); err != nil {
		ls.log.WithField("can't connect to database err", err.Error()).Error()
		os.Exit(1)
	}

	ls.restServer.Run()
}

// Stop ...
func (ls *loginServer) Stop() {

}
