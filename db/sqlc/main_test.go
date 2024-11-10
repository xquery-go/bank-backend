package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var (
	testQueries *Queries
	dbDriver    string
	dbUrl       string
)

// load .env before test
func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("cannot load env variables:", err)
	}

	// set global vars for test
	dbDriver = os.Getenv("POSTGRES_DRIVER")
	dbUrl = os.Getenv("POSTGRES_URL")
}

// test entry
func TestMain(m *testing.M) {

	conn, err := sql.Open(dbDriver, dbUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close()

	testQueries = New(conn)

	os.Exit(m.Run())

}
