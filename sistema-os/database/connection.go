package database

import (
	"context"
	"fmt"

	"sistema-os/config"

	"github.com/jackc/pgx/v5"
)

func Conectar(cfg config.Config) (*pgx.Conn, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar no PostgreSQL: %w", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		conn.Close(context.Background())
		return nil, fmt.Errorf("erro ao testar conexão com PostgreSQL: %w", err)
	}

	return conn, nil
}
