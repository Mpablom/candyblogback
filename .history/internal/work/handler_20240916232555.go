package work

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	repo *Repository
}

func RegisterRoutes(r *gin.Engine, db *mongo.Database) {
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
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	work, err := h.repo.GetWork(id)
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

	work.ID = primitive.NewObjectID()
	if err := h.repo.CreateWork(&work); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, work)
}

func (h *handler) UpdateWork(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedWork Work
	if err := c.ShouldBindJSON(&updatedWork); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedWork.ID = id

	if err := h.repo.UpdateWork(&updatedWork); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedWork)
}

func (h *handler) DeleteWork(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.repo.DeleteWork(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
