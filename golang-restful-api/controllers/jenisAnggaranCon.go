package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllJenisAnggaran(c *gin.Context) {
	var jenisAnggaran []models.JenisAnggaran

	if err := setup.DB.Find(&jenisAnggaran).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jenisAnggaran})
}
