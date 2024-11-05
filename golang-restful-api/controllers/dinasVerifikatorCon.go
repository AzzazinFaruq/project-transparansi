package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDinasVerifikator(c *gin.Context) {
	var dinasVerifikator []models.DinasVerifikator

	if err := setup.DB.Find(&dinasVerifikator).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dinasVerifikator})
}

func CreateDinasVerifikator(c *gin.Context) {
	var input struct {
		NamaDinasVerifikator string `json:"nama_dinas_verifikator" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dinasVerifikator := models.DinasVerifikator{
		NamaDinasVerifikator: input.NamaDinasVerifikator,
	}

	if err := setup.DB.Create(&dinasVerifikator).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func GetDinasVerifikatorByNama(c *gin.Context) {
	query := c.Query("nama_dinas_verifikator")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pencarian tidak boleh kosong"})
		return
	}

	var dinasVerifikator []models.DinasVerifikator

	if err := setup.DB.
		Where("nama_dinas_verifikator LIKE ?", "%"+query+"%").
		Find(&dinasVerifikator).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dinasVerifikator})
}

func DeleteDinasVerifikator(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak boleh kosong"})
		return
	}

	if err := setup.DB.Delete(&models.DinasVerifikator{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dinas Verifikator berhasil dihapus"})
}
