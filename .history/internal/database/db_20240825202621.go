package database

import (
	"log"

	"github.com/Mpablom/candyblogback/config"
	"github.com/Mpablom/candyblogback/internal/work"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	var(
		host = "localhost"
		user = "pablo"
		port = 5432
		password = "32899906"
		name = "candyblog"
	)
	var DB *gorm.DB  
 DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d 
 user=%s password=%s dbname=%s sslmode=disable", host, portInt, 
 user, password, name))
}
