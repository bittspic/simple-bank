package db

import (
	"database/sql"
	"log"
	"os"
	"simplebank/util"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	// var err error
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
