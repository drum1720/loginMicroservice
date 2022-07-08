package posgresql

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"loginMicroservice/app/pkg/configs"
	"time"
)

type DbConnectionPool struct {
	DbPool *pgxpool.Pool
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

	return &DbConnectionPool{DbPool: ConnectionPool}, nil
}

func (d *DbConnectionPool) Ping(ctx context.Context) error {
	dataSourceName := configs.GetEnvDefault("DSN", "")
	if dataSourceName == "" {
		return errors.New("no DATABASE env set")
	}

	if err := d.DbPool.Ping(ctx); err != nil {
		return err
	}
	return nil
}
