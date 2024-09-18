package main

import (
	"github.com/Mpablom/candyblogback/config"
	"github.com/Mpablom/candyblogback/internal/work"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectMongoDB()
	r := gin.Default()

	work.RegisterRoutes(r, config.GetMongoDB())

	r.Run(":8080")
}
