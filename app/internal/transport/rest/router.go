package rest

import (
	"context"
	"github.com/gorilla/mux"
	"loginMicroservice/app/internal/configs"
	"loginMicroservice/app/internal/datasource"
	"loginMicroservice/app/internal/logger"
	"loginMicroservice/app/internal/transport/rest/handlers"
	"loginMicroservice/app/internal/transport/rest/middlewares"
	"net/http"
)

type Router struct {
	ctx context.Context
	log logger.Logger
	cfg configs.Configure
}

// InitRouter ...
func InitRouter(
	ctx context.Context,
	log logger.Logger,
	cfg configs.Configure,
	dbSourcer datasource.DbSourcer,
) http.Handler {
	router := mux.NewRouter()

	adminRouter := router.PathPrefix("/admin/").Subrouter()
	adminRouter.Use(middlewares.NewAuth(cfg).Middleware) //authenticate
	adminRouter.Handle("/secretsequrdangerwork", handlers.NewSecretSequrDangerWorkHandler(ctx, dbSourcer, log)).Methods(http.MethodPost)

	userRouter := router.PathPrefix("/user/").Subrouter()
	userRouter.Handle("/register", handlers.NewRegistrationHandler(ctx, dbSourcer, log)).Methods(http.MethodPost)
	userRouter.Handle("/authorization", handlers.NewAuthorizeHandler(ctx, dbSourcer, log, cfg)).Methods(http.MethodPost)

	return router
}
