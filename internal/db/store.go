package db

import (
	postgresql "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/postgresql/sqlc"
	"github.com/jackc/pgx/v4/pgxpool"
)

//go:generate mockgen -source=./store.go -destination=./__mock__/store.go
type Store interface {
	postgresql.Querier
}

type storeImpl struct {
	postgresql.Querier
}

func NewStore(db *pgxpool.Pool) Store {
	return &storeImpl{
		Querier: postgresql.New(db),
	}
}
