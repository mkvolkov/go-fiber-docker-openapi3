package postgres

import (
	"context"
	"employees/pkg/config"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(cfg *config.Config) (*pgx.Conn, error) {
	pgURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
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
