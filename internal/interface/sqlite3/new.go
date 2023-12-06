package sqlite3

import (
	"database/sql"
)

type SQLite3Repository struct {
	DB *sql.DB
}

func New(db *sql.DB) *SQLite3Repository {
	return &SQLite3Repository{db}
}
