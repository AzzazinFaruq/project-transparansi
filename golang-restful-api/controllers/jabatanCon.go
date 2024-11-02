package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllJabatan(c *gin.Context) {
	var jabatan []models.Jabatan

	if err := setup.DB.Find(&jabatan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jabatan})
}

func CreateJabatan(c *gin.Context) {
	var input struct {
		Jabatan string `json:"jabatan" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jabatan := models.Jabatan{
		Jabatan: input.Jabatan,
	}

	if err := setup.DB.Create(&jabatan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func GetJabatanByNamaJabatan(c *gin.Context) {
	query := c.Query("jabatan")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pencarian tidak boleh kosong"})
		return
	}

	var jabatan []models.Jabatan

	if err := setup.DB.Where("jabatan LIKE ?", "%"+query+"%").Find(&jabatan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(jabatan) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Tidak ada jabatan yang ditemukan",
			"data":    []interface{}{},
		})
		return
	}

	// Format hasil pencarian
	formattedJabatan := make([]gin.H, len(jabatan))
	for i, jabatan := range jabatan {
		formattedJabatan[i] = gin.H{
			"id":       jabatan.Id,
			"jabatan": jabatan.Jabatan,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Jabatan ditemukan",
		"total":   len(jabatan),
		"data":    formattedJabatan,
	})
}

func DeleteJabatan(c *gin.Context) {
	id := c.Param("id")

	if err := setup.DB.Delete(&models.Jabatan{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Jabatan berhasil dihapus"})
}

