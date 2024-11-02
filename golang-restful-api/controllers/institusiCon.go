package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllInstitusi(c *gin.Context) {
	var institusi []models.Institusi

	if err := setup.DB.Find(&institusi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": institusi})
}

func CreateInstitusi(c *gin.Context) {
	var input struct {
		Institusi string `json:"institusi" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	

	institusi := models.Institusi{
		NamaInstitusi: input.Institusi,
	}

	if err := setup.DB.Create(&institusi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}
