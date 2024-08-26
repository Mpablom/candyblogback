package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	var (
		host     = "localhost"
		user     = "pablo"
		port     = 5432
		password = "32899906"
		name     = "candyblog"
	)
	var DB *gorm.DB
	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name))
	if err != nil {
		log.Fatalf("Error in connect the DB %v", err)
		return nil
	}
	if err := DB.DB().Ping(); err != nil {
		log.Fatalln("Error in make ping the DB " + err.Error())
		return nil
	}
	if DB.Error != nil {
		log.Fatalln("Any Error in connect the DB " + err.Error())
		return nil
	}
	log.Println("DB connected")
	return DB
}
