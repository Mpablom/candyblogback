package work

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	repo *Repository
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := newRepository(db)
	h := &handler{repo: repo}
	r.GET("/", h.HandleRoot)
	r.GET("/works", h.GetAllWorks)
	r.GET("/works/:id", h.GetWork)
	r.POST("/works", h.CreateWork)
	r.PUT("/works/:id", h.UpdateWork)
	r.DELETE("/works/:id", h.DeleteWork)
}
func (h *handler) HandleRoot(c *gin.Context) {
	c.String(http.StatusOK, "CandyBlog está corriendo!")
}
func (h *handler) GetAllWorks(c *gin.Context) {
	works, err := h.repo.GetAllWorks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, works)
}

func (h *handler) GetWork(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	work, err := h.repo.GetWork(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, work)
}

func (h *handler) CreateWork(c *gin.Context) {
	var work Work
	if err := c.ShouldBindJSON(&work); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.CreateWork(&work); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for i := range work.Gallery {
		work.Gallery[i].WorkID = work.ID
	}
	for _, gallery := range work.Gallery {
		if err := h.repo.CreateGallery(&gallery); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusCreated, work)
}

func (h *handler) UpdateWork(c *gin.Context) {
	var work Work
	if err := c.ShouldBindJSON(&work); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.UpdateWork(&work); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, work)
}

func (h *handler) DeleteWork(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.repo.DeleteWork(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
