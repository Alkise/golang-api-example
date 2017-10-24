package database

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// Pool Pool of database connection
	Pool *sql.DB
)

// InitDB Initialize database connection
func InitDB(provider string, databaseName string) (*sql.DB, error) {
	db, err := sql.Open(provider, databaseName)

	if err != nil {
		return nil, err
	}

	if db == nil {
		return nil, errors.New("db is nil")
	}
	return db, nil
}
