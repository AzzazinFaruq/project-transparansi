package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAduan(c *gin.Context) {
	var Aduan []models.Aduan

	if err := setup.DB.
		Preload("Program").
		Preload("User.Jabatan").
		Preload("User.Role").
		Find(&Aduan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Aduan})
}
