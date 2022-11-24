package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	db "github.com/nina/db/sqlc/project_info/gen"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:nina@localhost:5435/nina_app?sslmode=disable"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}
	testQueries = db.New(testDB)
	os.Exit(m.Run())
}
