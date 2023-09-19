package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/gaomengnan/go-micro/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:123456@localhost/b_micro?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal(err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
