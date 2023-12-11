package config

import (
	"os"
)

func LoadSQLite3DBName() string {
	dbName, ok := os.LookupEnv("SQLITE3_DBNAME")

	if ok {
		return dbName
	}

	panic("SQLITE3_DBNAME IS NOT SET")
}
