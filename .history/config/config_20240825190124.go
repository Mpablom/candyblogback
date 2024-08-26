package config

import (
	"os"
)

func GetDatabaseUrl() string {
	return os.Getenv("DATABASE_URL")
}
