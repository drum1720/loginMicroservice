package posgresql

import (
	"context"
	"github.com/pkg/errors"
	"loginMicroservice/app/internal/core"
	"loginMicroservice/app/internal/security"
	"loginMicroservice/app/internal/transport/rest/server/response"
	"net/http"
	"time"
)

func (d DbConnectionPool) UserExist(ctx context.Context, user core.User) (bool, *response.Error) {
	query := "SELECT login from auf WHERE login=$1"
	res, err := d.DbPool.Exec(ctx, query, user.User)
	if err != nil {
		return false, response.NewError(err, http.StatusInternalServerError)
	}

	dd := string(res)
	if dd != "SELECT 0" {
		return true, response.NewError(errors.Errorf("user %s exist", user.User), http.StatusInternalServerError)
	}

	return false, nil
}

func (d DbConnectionPool) InsertUser(ctx context.Context, user core.User) *response.Error {
	pass := security.PasswordEncryption(user.Password)
	if pass == nil {
		return response.NewError(errors.New("server error"), http.StatusInternalServerError)
	}

	query := "INSERT INTO auf (login, pass, last_visit) VALUES ($1,$2,$3)"
	_, err := d.DbPool.Exec(ctx, query, user.User, pass, time.Now())
	if err != nil {
		return response.NewError(err, http.StatusInternalServerError)
	}

	return nil
}

func (d DbConnectionPool) UserValid(ctx context.Context, user *core.User) (bool, *response.Error) {
	requestPass := user.Password
	query := "SELECT login,pass,last_visit from auf WHERE login=$1"

	res := d.DbPool.QueryRow(ctx, query, user.User)
	if res == nil {
		return false, response.NewError(errors.Errorf("user %s not exist", user.User), http.StatusBadRequest)
	}

	if err := res.Scan(&user.User, &user.Password, &user.LastVisit); err != nil {
		return false, response.NewError(err, http.StatusInternalServerError)
	}

	if !security.PassCorrect(requestPass, user.Password) {
		return false, response.NewError(errors.Errorf("password for user: '%s' not valid", user.User), http.StatusBadRequest)
	}

	return true, nil
}
