package datasource

import (
	"context"
	"loginMicroservice/app/internal/core"
	error2 "loginMicroservice/app/internal/transport/rest/error"
)

type DbSourcer interface {
	Ping(context.Context) error
	UserExist(context.Context, core.User) (bool, *error2.Error)
	InsertUser(context.Context, core.User) *error2.Error
	UserValid(context.Context, *core.User) (bool, *error2.Error)
}
