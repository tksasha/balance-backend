package app

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tksasha/balance/config"
)

func OpenDB(config *config.Config) *sql.DB {
	db, err := sql.Open("sqlite3", config.SQLite3DBName)
	if err != nil {
		panic(err)
	}

	return db
}
