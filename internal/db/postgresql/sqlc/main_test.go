package postgresql

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var queries *Queries

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("something went wrong %v", err)
	}
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_CONNECTION_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	queries = New(conn)

	os.Exit(m.Run())
}
