package main

import (
	"candyblogback/internal/database"
	"candyblogback/internal/work"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDatabase()
	r := gin.Default()
	work.RegisterRoutes(r, db)
	r.Run()
}
