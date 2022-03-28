package postgresql

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/DSC-UNSRI/gdsc-website-backend/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pashagolub/pgxmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

// Use *queries (struct) for tests only, DON'T USE STRUCT FOR YOUR CODE, USE INTERFACE INSTEAD !!!!!!!!!
var querier *Queries

func downMigration() {
	out, err := exec.Command("make", "-C", "../../../../", "migrate-down-test").CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		log.Fatalln("something went wrong when dropping db")
	}
}

func upMigration() {
	out, err := exec.Command("make", "-C", "../../../../", "migrate-up-test").CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		log.Fatalln("something went wrong when up db")
	}
}

func useMockDB(t *testing.T) pgxmock.PgxConnIface {
	mock, err := pgxmock.NewConn()
	require.NoError(t, err)
	return mock
}

func TestMain(m *testing.M) {
	upMigration()
	config.New("../../../../.env")
	dbConfig, err := pgxpool.ParseConfig(viper.GetString("DB_CONNECTION_TEST_URL"))
	if err != nil {
		log.Fatalf("something went wrong %v", err)
	}
	conn, err := pgxpool.ConnectConfig(context.Background(), dbConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	querier = New(conn)

	exitCode := m.Run()
	downMigration()

	os.Exit(exitCode)
}
