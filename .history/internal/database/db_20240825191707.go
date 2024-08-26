package database

import (
	"log"

	"github.com/Mpablom/candyblogback/config"
	"github.com/Mpablom/candyblogback/internal/work"

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
