package application

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"loginMicroservice/app/internal/core"
	"loginMicroservice/app/internal/datasource/posgresql"
	"loginMicroservice/app/internal/transport/grpc/server"
	"loginMicroservice/app/internal/transport/grpc/server/handlers"
	"net/http"
	"os"
)

type loginServer struct {
	logger           *logrus.Logger
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
	var err error

	ls.cfg, err = core.InitCfg()
	if err != nil {
		ls.logger.Errorf("config load err: %s", err)
	}

	ls.dbConnectionPool, err = posgresql.NewDbConnectionPool(ls.cfg.GetDsnPG())
	if err != nil {
		ls.logger.Errorf("can't connect to database err: %s", err.Error())
		os.Exit(1)
	}

	ls.logger = logrus.New()
	ls.ctx = context.Background()

	router := mux.NewRouter()
	router.Handle("/authorization", handlers.NewRegistrationHandler(ls.ctx, ls.dbConnectionPool, ls.logger)).Methods(http.MethodPost)

	ls.restServer = server.NewRestServer(ls.ctx, ls.logger, ls.cfg.GetUrl(), router)
}

// Run ...
func (ls *loginServer) Run() {
	if err := ls.dbConnectionPool.Ping(ls.ctx); err != nil {
		ls.logger.Errorf("can't connect to database err: %s", err.Error())
		os.Exit(1)
	}

	ls.restServer.Run()
}

// Stop ...
func (ls *loginServer) Stop() {

}
