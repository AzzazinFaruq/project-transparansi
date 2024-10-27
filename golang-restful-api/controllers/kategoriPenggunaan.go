package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllKategoriPenggunaan(c *gin.Context) {
	var kategoriPenggunaan []models.KategoriPenggunaan

	if err := setup.DB.Find(&kategoriPenggunaan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": kategoriPenggunaan})
}
