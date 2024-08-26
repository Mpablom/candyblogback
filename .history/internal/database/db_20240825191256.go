package database

import (
	"candyblogback/config"
	"candyblogback/internal/work"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := config.GetDatabaseUrl()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	db.AutoMigrate(&work.Work{}, &work.Gallery{})
	return db
}
