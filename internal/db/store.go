package db

import (
	postgresql "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/postgresql/sqlc"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	postgresql.Querier
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		Querier: postgresql.New(db),
	}
}
