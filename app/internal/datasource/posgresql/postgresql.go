package posgresql

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

const (
	maxConn         = 50
	minConn         = 10
	maxConnLifeTime = 15 * time.Minute
)

type DbConnectionPool struct {
	dbPool *pgxpool.Pool
}

//NewDbConnectionPool ...
func NewDbConnectionPool(dataSourceName string) (*DbConnectionPool, error) {
	connectionPool, err := pgxpool.Connect(context.Background(), dataSourceName)
	if err != nil {
		return nil, err
	}

	connectionPool.Config().MaxConns = maxConn
	connectionPool.Config().MinConns = minConn
	connectionPool.Config().MaxConnLifetime = maxConnLifeTime

	return &DbConnectionPool{dbPool: connectionPool}, nil
}

// Ping ...
func (d *DbConnectionPool) Ping(ctx context.Context) error {
	if err := d.dbPool.Ping(ctx); err != nil {
		return err
	}
	return nil
}
