package storage

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	DB  *pgx.Conn
	Log *slog.Logger
	Ctx context.Context
}

func InitStorage(s Storage) *Storage {
	return &Storage{
		DB:  s.DB,
		Log: s.Log,
		Ctx: s.Ctx,
	}
}
