package postgresql_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/DSC-UNSRI/gdsc-website-backend/config"
	postgresql "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/postgresql/sqlc"
	"github.com/jackc/pgx/v4/pgxpool"
)

var querier postgresql.Querier

func TestMain(m *testing.M) {
	config := config.New("../../../../.env")
	dbConfig, err := pgxpool.ParseConfig(config.PostgresDSN)
	if err != nil {
		log.Fatalf("something went wrong %v", err)
	}
	conn, err := pgxpool.ConnectConfig(context.Background(), dbConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	querier = postgresql.New(conn)

	os.Exit(m.Run())
}
