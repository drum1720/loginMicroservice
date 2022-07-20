package posgresql

import (
	"context"
	"github.com/pkg/errors"
	"loginMicroservice/app/internal/security"
	error2 "loginMicroservice/app/internal/transport/rest/error"
	"net/http"
	"time"
)

func (d DbConnectionPool) UserExist(ctx context.Context, user string) (bool, *error2.Error) {
	query := "SELECT login from auf WHERE login=$1"

	execResult, err := d.dbPool.Exec(ctx, query, user)
	if err != nil {
		return false, error2.NewError(err, http.StatusInternalServerError)
	}

	if string(execResult) != "SELECT 0" {
		return true, error2.NewError(errors.Errorf("user %s exist", user), http.StatusInternalServerError)
	}

	return false, nil
}

func (d DbConnectionPool) InsertUser(ctx context.Context, user, pass string) *error2.Error {
	passHash := security.PassEncryption(pass)
	if passHash == nil {
		return error2.NewError(errors.New("server error"), http.StatusInternalServerError)
	}

	query := "INSERT INTO auf (login, pass, last_visit) VALUES ($1,$2,$3)"
	_, err := d.dbPool.Exec(ctx, query, user, passHash, time.Now())
	if err != nil {
		return error2.NewError(err, http.StatusInternalServerError)
	}

	return nil
}

func (d DbConnectionPool) UserValid(ctx context.Context, user, pass string) (bool, *error2.Error) {
	var id int
	requestPass := ""
	query := "SELECT pass, id from auf WHERE login=$1"

	queryRow := d.dbPool.QueryRow(ctx, query, user)
	if queryRow == nil {
		return false, error2.NewError(errors.Errorf("user %s not exist", user), http.StatusBadRequest)
	}

	if err := queryRow.Scan(&requestPass, &id); err != nil {
		return false, error2.NewError(err, http.StatusInternalServerError)
	}

	if !security.PassCorrect(pass, requestPass) {
		return false, error2.NewError(errors.Errorf("password for user: '%s' not valid", user), http.StatusBadRequest)
	}

	query = "UPDATE auf SET last_visit = $1 WHERE id = $2"
	if _, err := d.dbPool.Exec(ctx, query, time.Now(), id); err != nil {
		return false, error2.NewError(err, http.StatusInternalServerError)
	}

	return true, nil
}

func (d DbConnectionPool) ClearTable(ctx context.Context) *error2.Error {
	query := "DELETE FROM auf"
	if _, err := d.dbPool.Exec(ctx, query); err != nil {
		return error2.NewError(err, http.StatusInternalServerError)
	}

	return nil
}
