package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tksasha/balance/config"
)

func Open(c *config.Config) *sql.DB {
	db, err := sql.Open("sqlite3", c.SQLite3Config.DB)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		log.Fatal(err)
	}

	return db
}
