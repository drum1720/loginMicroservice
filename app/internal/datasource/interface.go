package datasource

import (
	"context"
	error2 "loginMicroservice/app/internal/transport/rest/error"
)

type (
	DbSourcer interface {
		Ping(context.Context) error
		UserExist(ctx context.Context, user string) (bool, *error2.Error)
		InsertUser(ctx context.Context, user, pass string) *error2.Error
		UserValid(ctx context.Context, user, pass string) (bool, *error2.Error)
	}
)
