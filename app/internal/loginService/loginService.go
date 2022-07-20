package loginService

import (
	"context"
	"loginMicroservice/app/internal/configs"
	"loginMicroservice/app/internal/datasource"
	"loginMicroservice/app/internal/datasource/posgresql"
	"loginMicroservice/app/internal/logger"
	"loginMicroservice/app/internal/logger/logrus"
	"loginMicroservice/app/internal/transport"
	"loginMicroservice/app/internal/transport/rest"
	"net/http"
	"os"
	"time"
)

type LoginServer struct {
	ctx        context.Context
	log        logger.Logger
	cfg        configs.Configure
	dbSourcer  datasource.DbSourcer
	router     http.Handler
	restServer transport.Server
	cancel     func()
}

// NewLoginServer ...
func NewLoginServer() *LoginServer {
	return &LoginServer{}
}

// Init ...
func (ls *LoginServer) Init() {
	ls.log = logrus.NewLogger()
	ls.ctx, ls.cancel = context.WithCancel(context.Background())

	var err error
	ls.cfg, err = configs.InitCfg()
	if err != nil {
		ls.log.WithField("config err", err).Error()
		os.Exit(1)
	}

	ls.dbSourcer, err = posgresql.NewDbConnectionPool(ls.cfg.GetDsnPG())
	if err != nil {
		ls.log.WithField("can't connect to database err", err.Error()).Error()
		os.Exit(1)
	}
	ls.router = rest.InitRouter(ls.ctx, ls.log, ls.cfg, ls.dbSourcer)
	ls.restServer = rest.NewRestServer(ls.ctx, ls.log, ls.cfg, ls.router)
}

// Run ...
func (ls *LoginServer) Run() {
	if err := ls.dbSourcer.Ping(ls.ctx); err != nil {
		ls.log.WithField("can't connect to database err", err.Error()).Error()
		os.Exit(1)
	}

	go ls.restServer.Run(ls.cancel)

	select {
	case <-ls.ctx.Done():
		ls.restServer.Shutdown()
		time.Sleep(time.Second * 5)
		os.Exit(1)
	}
}
