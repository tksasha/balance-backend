package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Open() (db *sql.DB) {
	var err error

	if db, err = sql.Open("sqlite3", "db/development.sqlite3"); err != nil {
		log.Fatal(err)
	}

	return
}

func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
