package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/dangquangdon/simplebank/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, errConf := utils.LoadConfig("../../environments", "testing")
	if errConf != nil {
		return
	}

	var err error
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDb)
	os.Exit(m.Run())
}
