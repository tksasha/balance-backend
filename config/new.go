package config

import (
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	SQLite3Config SQLite3Config
}

func New() *Config {
	setTimeZone()

	return &Config{
		SQLite3Config: loadSQLite3Config(),
	}
}
