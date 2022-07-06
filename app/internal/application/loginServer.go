package application

import (
	"context"
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

	//добавляем хендлеры, они будут добавлены в роутер
	handlersMap := make(map[string]http.Handler)
	handlersMap["/authorization"] = handlers.NewAuthorization(ls.ctx, ls.dbConnectionPool, ls.logger)

	//создание сервера,регистрация хендлеров в роутер
	ls.restServer = server.NewRestServer(ls.ctx, ls.logger, ls.cfg.GetUrl(), handlersMap)
}

// Run ...
func (ls *loginServer) Run() {
	if err := ls.dbConnectionPool.Ping(ls.ctx); err != nil {
		ls.logger.Errorf("can't connect to database err: %s", err.Error())
		os.Exit(1)
	}

	ls.restServer.Run()
	// старт сервера и назначение роутинга

	//...
}

// Stop ...
func (ls *loginServer) Stop() {

}
