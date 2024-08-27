package main

import (
	"github.com/Mpablom/candyblogback/config"
	"github.com/Mpablom/candyblogback/internal/work"
	"github.com/gin-gonic/gin"
)

func main() {
	dbConfig := config.Configure("./", "postgresql")
	config.DB = dbConfig.InitPostgresDB()
	config.DB.AutoMigrate(&work.Work{}, &work.Gallery{})

	r := gin.Default()

	work.RegisterRoutes(r, config.DB)

	r.Run(":8080")
}
