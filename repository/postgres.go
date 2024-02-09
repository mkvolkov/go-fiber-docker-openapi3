package repository

import (
	"context"
	"employees/config"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type PgCfg struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func ConnectDB(cfg *config.Postgres) (*pgx.Conn, error) {
	pgURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	conn, err := pgx.Connect(context.Background(), pgURL)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
