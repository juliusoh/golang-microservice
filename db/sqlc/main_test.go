package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_"github.com/lib/pq" // import the postgres driver _ blank identifier
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
)

// Global variable, use it in other test files
// DB Connection to create Queries Object
var testQueries *Queries
var testDB *sql.DB
// Main entry point of all unit tests inside this package db
func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	// sql.Open() returns a connection object and an error

	// if error is not null is same thing as if (error)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	// Create Queries Object
	// use the connection to create a Queries Object
	testQueries = New(testDB)

	// m.Run to start running the unit test, it will return an exit code
	// Report the exit code to the OS
	os.Exit(m.Run())
}