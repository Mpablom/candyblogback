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
		work.Gallery[i].ID = 0
	}

	c.JSON(http.StatusCreated, work)
}

func (h *handler) UpdateWork(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Obtener el work existente
	existingWork, err := h.repo.GetWork(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Work not found"})
		return
	}

	// Obtener los datos del work actualizado desde la solicitud
	var updatedWork Work
	if err := c.ShouldBindJSON(&updatedWork); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Iniciar una transacción
	tx := h.repo.db.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed"})
		}
	}()

	// Eliminar las galerías existentes
	if err := tx.Where("work_id = ?", existingWork.ID).Delete(&Gallery{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Actualizar el work
	existingWork.Image = updatedWork.Image
	existingWork.Title = updatedWork.Title
	existingWork.Description = updatedWork.Description
	if err := tx.Save(existingWork).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Crear nuevas galerías
	for _, gallery := range updatedWork.Gallery {
		gallery.WorkID = existingWork.ID
		if err := tx.Create(&gallery).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Confirmar la transacción
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Cargar el work actualizado con las galerías
	if err := h.repo.db.Preload("Gallery").First(&existingWork, existingWork.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingWork)
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
