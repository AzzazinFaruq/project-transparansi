package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllLog(c *gin.Context) {
	var log []models.Log

	if err := setup.DB.Preload("User").Find(&log).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Buat slice baru untuk menyimpan data log yang telah diformat
	formattedLog := make([]gin.H, len(log))

	for i, log := range log {
		formattedLog[i] = gin.H{
			"id":         log.Id,
			"user_id":    log.UserId,
			"aktivitas":  log.Aktivitas,
			"status":     log.Status,
			"created_at": log.FormattedCreatedAt(),
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": formattedLog})
}
