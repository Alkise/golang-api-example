package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// Pool Pool of database connection
	Pool *sql.DB
)

// InitDB Initialize database connection
func InitDB(provider string, databaseName string) *sql.DB {
	db, err := sql.Open(provider, databaseName)

	panicOnErr(err)

	if db == nil {
		panic("db is nil")
	}
	return db
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
