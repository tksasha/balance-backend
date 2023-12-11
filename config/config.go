package config

type Config struct {
	SQLite3DBName string
}

func NewConfig() *Config {
	return &Config{
		SQLite3DBName: LoadSQLite3DBName(),
	}
}
