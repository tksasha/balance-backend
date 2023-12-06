package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	SQLite3Config SQLite3Config
}

func New() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	setTimeZone()

	return &Config{
		SQLite3Config: loadSQLite3Config(),
	}
}
