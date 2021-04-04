package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver string = "postgres"
	dbSource string = "postgresql://postgres:postgres@127.0.0.1:5432/simplebank?sslmode=disable"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDb)
	os.Exit(m.Run())
}
