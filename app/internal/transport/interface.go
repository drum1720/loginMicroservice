package transport

import (
	"context"
	"loginMicroservice/app/internal/configs"
	"loginMicroservice/app/internal/datasource"
	"loginMicroservice/app/internal/logger"
	"net/http"
)

type Server interface {
	Run(cancel func())
	Shutdown()
}

type Router interface {
	InitRouter(context.Context, logger.Logger, configs.Configure, datasource.DbSourcer) http.Handler
}
