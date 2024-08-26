package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: %v", err)
	} else {
		log.Println("Loaded .env file")
	}
}
func GetDatabaseURL() string {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("DATABASE_URL not set")
	}
	return url
}

func GetJPAConfig() (ddlAuto string, openInView string) {
	ddlAuto = os.Getenv("JPA_DDL_AUTO")
	openInView = os.Getenv("JPA_OPEN_IN_VIEW")
	return
}
