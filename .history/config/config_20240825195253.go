package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
func GetDatabaseUrl() string {
	return os.Getenv("DATABASE_URL")
}

func GetJPAConfig() (ddlAuto string, openInView string) {
	ddlAuto = os.Getenv("JPA_DDL_AUTO")
	openInView = os.Getenv("JPA_OPEN_IN_VIEW")
	return
}
