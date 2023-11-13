package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "db/development.sqlite3")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
