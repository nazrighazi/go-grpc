package dataaccess

import (
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type DataAccessDto struct {
	Db  *pgx.Conn
	Log *slog.Logger
}

func NewDataAccess(req DataAccessDto) *Queries {
	queries := New(req.Db)
	return queries
}