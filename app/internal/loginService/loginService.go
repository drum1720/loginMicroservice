package loginService

import (
	"context"
	"github.com/gorilla/mux"
	"loginMicroservice/app/internal/configs"
	"loginMicroservice/app/internal/datasource/posgresql"
	"loginMicroservice/app/internal/logger"
	"loginMicroservice/app/internal/logger/logrus"
	"loginMicroservice/app/internal/transport/rest/server"
	"loginMicroservice/app/internal/transport/rest/server/handlers"
	"net/http"
	"os"
)

type LoginServer struct {
	log              logger.Logger
	ctx              context.Context
	cfg              *configs.Cfg
	dbConnectionPool *posgresql.DbConnectionPool
	restServer       *server.RestServer
}

// NewLoginServer ...
func NewLoginServer() *LoginServer {
	return &LoginServer{}
}

// Init ...
func (ls *LoginServer) Init() {
	ls.log = logrus.NewLogger()
	ls.ctx = context.Background()

	var err error
	ls.cfg, err = configs.InitCfg()
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
func (ls *LoginServer) Run() {
	if err := ls.dbConnectionPool.Ping(ls.ctx); err != nil {
		ls.log.WithField("can't connect to database err", err.Error()).Error()
		os.Exit(1)
	}

	ls.restServer.Run()
}

// Stop ...
func (ls *LoginServer) Stop() {

}

// Restart ...
func (ls *LoginServer) Restart() {

}
