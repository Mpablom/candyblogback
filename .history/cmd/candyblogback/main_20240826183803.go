package main

import (
	"github.com/Mpablom/candyblogback/config"
	"github.com/Mpablom/candyblogback/internal/work"
)

func main() {
	dbConfig := config.Configure("./", "postgresql")
	config.DB = dbConfig.InitPostgresDB()
	config.DB.AutoMigrate(&work.Work{}, &work.Gallery{})
}
