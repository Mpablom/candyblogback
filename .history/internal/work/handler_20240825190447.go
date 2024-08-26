package work

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{db: db}

	r.GET("/works", h.GetAllWorks)
	r.GET("/works/:id", h.GetWork)
	r.POST("/works", h.CreateWork)
	r.PUT("/works/:id", h.UpdateWork)
	r.DELETE("/works/:id", h.DeleteWork)
}
