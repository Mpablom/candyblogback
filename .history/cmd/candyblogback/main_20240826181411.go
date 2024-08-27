package main

import (
	"github.com/Mpablom/candyblogback/config"
	"github.com/Mpablom/candyblogback/internal/work"
)

func main() {
	dbConfig := config.Configure("./", "postgres")
	config.DB = dbConfig.InitPostgresDB()
	config.DB.AutoMigrate(&work.Work{})
}
