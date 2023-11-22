package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Open() (db *sql.DB) {
	var err error

	if db, err = sql.Open("sqlite3", dbname()); err != nil {
		log.Fatal(err)
	}

	return
}

func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

func dbname() string {
	switch os.Getenv("GOENV") {
		case "production":
			return "db/production.sqlite3"
		case "test":
			return "db/test.sqlite3"
		default:
			return "db/development.sqlite3"
	}
}
