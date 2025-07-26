package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Config struct {
	Context          context.Context
	ConnectingString string
}

func DatabaseInit(cfg *Config) (*pgx.Conn, error) {
	db, err := pgx.Connect(cfg.Context, cfg.ConnectingString)
	if err != nil {
		return nil, err
	} else {
		return db, nil
	}
}
