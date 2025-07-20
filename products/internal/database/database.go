package database

import (
	"context"
	"fmt"
	"go_grpc/products/internal/config"

	"github.com/jackc/pgx/v5"
)

func NewDatabase(ctx context.Context,conf *config.Config) (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(ctx, fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", conf.Db.User, conf.Db.Password, conf.Db.Host, conf.Db.Port, conf.Db.DBName, conf.Db.SSLMode))
	if err != nil {
		return nil, err
	}


	return conn, nil
}