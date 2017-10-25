package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDB Initialize database connection
func InitDB(driver string, dataSourceName string) (*sql.DB, error) {
	var err error

	db, err = sql.Open(driver, dataSourceName)
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}
