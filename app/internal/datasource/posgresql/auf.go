package posgresql

import (
	"context"
	"github.com/pkg/errors"
	"loginMicroservice/app/internal/core"
	"loginMicroservice/app/internal/transport/grpc/server/response"
	"net/http"
	"time"
)

func (d DbConnectionPool) ExistUser(ctx context.Context, user core.User) bool {
	res, err := d.DbPool.Exec(ctx, "SELECT login from auf WHERE login=$1", user.User)
	if err != nil {
		return true
	}

	dd := string(res)
	if dd != "SELECT 0" {
		return true
	}

	return false
}

func (d DbConnectionPool) InsertUser(ctx context.Context, user core.User) *response.ResponseErr {
	_, err := d.DbPool.Exec(
		ctx,
		"INSERT INTO auf (login, pass, last_visit) VALUES ($1,$2,$3)",
		user.User, user.Password,
		time.Now())
	if err != nil {
		return response.NewResponseErr(err, http.StatusInternalServerError)
	}

	return nil
}

func (d DbConnectionPool) GetUser(ctx context.Context, user *core.User) *response.ResponseErr {
	res := d.DbPool.QueryRow(ctx, "SELECT login,pass,last_visit from auf WHERE login=$1", user.User)
	if res == nil {
		return response.NewResponseErr(errors.Errorf("user %s not exist", user.User), http.StatusBadRequest)
	}

	if err := res.Scan(user.User, user.Password, user.LastVisit); err != nil {
		return response.NewResponseErr(err, http.StatusInternalServerError)
	}

	return nil
}
