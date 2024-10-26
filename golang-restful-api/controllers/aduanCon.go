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

func CreateAduan(c *gin.Context) {
	var input struct {
		ProgramId int64  `json:"program_id" binding:"required"`
		UserId    int64  `json:"user_id" binding:"required"`
		Keluhan   string `json:"keluhan" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := setup.DB.Begin()

	aduan := models.Aduan{
		ProgramId: input.ProgramId,
		UserId:    input.UserId,
		Keluhan:   input.Keluhan,
		Status:    "Menunggu",
	}

	if err := tx.Create(&aduan).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat aduan"})
		return
	}

	newLog := models.Log{
		UserId:    input.UserId,
		Aktivitas: "Pengajuan Keluhan",
		Status:    "Menunggu",
	}
	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat log aktivitas"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}
	
	tx.Preload("Program").
		Preload("User.Jabatan").
		Preload("User.Role").
		First(&aduan, aduan.Id)

	c.JSON(http.StatusCreated, gin.H{"data": aduan})
}

func CountAduan(c *gin.Context){
	var count_aduan int64
	var count_aduan_disetujui int64
	var count_aduan_menunggu int64
	if err := setup.DB.Model(&models.Aduan{}).Count(&count_aduan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := setup.DB.Model(&models.Aduan{}).Where("status = ?", "Menunggu").Count(&count_aduan_menunggu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := setup.DB.Model(&models.Aduan{}).Where("status = ?", "Disetujui").Count(&count_aduan_disetujui).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": count_aduan,"disetujui": count_aduan_disetujui, "menunggu":count_aduan_menunggu})
}
