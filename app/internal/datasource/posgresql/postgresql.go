package posgresql

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"loginMicroservice/app/pkg/configs"
	"time"
)

type DbConnectionPool struct {
	dbPool *pgxpool.Pool
}

//NewDbConnectionPool ...
func NewDbConnectionPool(dataSourceName string) (*DbConnectionPool, error) {
	ConnectionPool, err := pgxpool.Connect(context.Background(), dataSourceName)
	if err != nil {
		return nil, err
	}

	ConnectionPool.Config().MaxConns = int32(configs.GetEnvAsInt("DB_MAX_OPEN_CONNECTION", 50))
	ConnectionPool.Config().MinConns = int32(configs.GetEnvAsInt("DB_MAX_OPEN_CONNECTION", 10))
	ConnectionPool.Config().MaxConnLifetime = time.Minute * 15

	return &DbConnectionPool{dbPool: ConnectionPool}, nil
}

func (d *DbConnectionPool) Ping(ctx context.Context) (err error) {
	dataSourceName := configs.GetEnvDefault("DATABASE", "")
	if dataSourceName == "" {
		return errors.New("no DATABASE env set")
	}

	if err := d.dbPool.Ping(ctx); err != nil {
		return err
	}
	return nil
}
