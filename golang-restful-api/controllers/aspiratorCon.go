package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAspirator(c *gin.Context) {
	var aspirator []models.Aspirator

	if err := setup.DB.Find(&aspirator).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": aspirator})
}

func CreateAspirator(c *gin.Context) {
	var input struct {
		NamaAspirator string `json:"nama_aspirator" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aspirator := models.Aspirator{
		NamaAspirator: input.NamaAspirator,
	}

	if err := setup.DB.Create(&aspirator).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func GetAspiratorByNamaAspirator(c *gin.Context) {
	query := c.Query("nama_aspirator")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pencarian tidak boleh kosong"})
		return
	}

	var aspirator []models.Aspirator

	if err := setup.DB.
		Where("nama_aspirator LIKE ?", "%"+query+"%").
		Find(&aspirator).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": aspirator})
}

func DeleteAspirator(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak boleh kosong"})
		return
	}

	if err := setup.DB.Delete(&models.Aspirator{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aspirator berhasil dihapus"})
}
