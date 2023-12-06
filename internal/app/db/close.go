package db

import (
	"database/sql"
	"log"
)

func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
