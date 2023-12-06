package config

import (
	"os"
)

type SQLite3Config struct {
	DB string
}

func loadSQLite3Config() SQLite3Config {
	return SQLite3Config{os.Getenv("DB")}
}
