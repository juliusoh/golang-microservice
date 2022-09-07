package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_"github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
)

// Global variable, use it in other test files
// DB Connection to create Queries Object
var testQueries *Queries

// Main entry point of all unit tests inside this package db
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}