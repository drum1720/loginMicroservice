package posgresql

import (
	"context"
	"github.com/pkg/errors"
	"loginMicroservice/app/internal/core"
	"loginMicroservice/app/internal/security"
	error2 "loginMicroservice/app/internal/transport/rest/error"
	"net/http"
	"time"
)

func (d DbConnectionPool) UserExist(ctx context.Context, user core.User) (bool, *error2.Error) {
	query := "SELECT login from auf WHERE login=$1"

	execResult, err := d.dbPool.Exec(ctx, query, user.User)
	if err != nil {
		return false, error2.NewError(err, http.StatusInternalServerError)
	}

	if string(execResult) != "SELECT 0" {
		return true, error2.NewError(errors.Errorf("user %s exist", user.User), http.StatusInternalServerError)
	}

	return false, nil
}

func (d DbConnectionPool) InsertUser(ctx context.Context, user core.User) *error2.Error {
	pass := security.PasswordEncryption(user.Password)
	if pass == nil {
		return error2.NewError(errors.New("server error"), http.StatusInternalServerError)
	}

	query := "INSERT INTO auf (login, pass, last_visit) VALUES ($1,$2,$3)"
	_, err := d.dbPool.Exec(ctx, query, user.User, pass, time.Now())
	if err != nil {
		return error2.NewError(err, http.StatusInternalServerError)
	}

	return nil
}

func (d DbConnectionPool) UserValid(ctx context.Context, user *core.User) (bool, *error2.Error) {
	requestPass := user.Password
	query := "SELECT login,pass,last_visit from auf WHERE login=$1"

	queryRow := d.dbPool.QueryRow(ctx, query, user.User)
	if queryRow == nil {
		return false, error2.NewError(errors.Errorf("user %s not exist", user.User), http.StatusBadRequest)
	}

	if err := queryRow.Scan(&user.User, &user.Password, &user.LastVisit); err != nil {
		return false, error2.NewError(err, http.StatusInternalServerError)
	}

	if !security.PassCorrect(requestPass, user.Password) {
		return false, error2.NewError(errors.Errorf("password for user: '%s' not valid", user.User), http.StatusBadRequest)
	}

	return true, nil
}
