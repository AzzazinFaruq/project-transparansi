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

	formattedAduans := make([]gin.H, len(Aduan))

	for i, aduan := range Aduan {
		formattedAduans[i] = gin.H{
			"id":         aduan.Id,
			"program_id": aduan.ProgramId,
			"program":    aduan.Program,
			"user_id":    aduan.UserId,
			"user":       aduan.User,
			"keluhan":    aduan.Keluhan,
			"status":     aduan.Status,
			"created_at": aduan.CreatedAt.Format("02-01-2006"), // Format d-m-y
			"updated_at": aduan.UpdatedAt.Format("02-01-2006"), // Format d-m-y
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": formattedAduans})
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

func GetAduanPerBulan(c *gin.Context) {
	tahun := c.Query("tahun")
	bulan := c.Query("bulan")

	if tahun == "" || bulan == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter tahun dan bulan diperlukan"})
		return
	}

	startDate := tahun + "-" + bulan + "-01"
	endDate := tahun + "-" + bulan + "-31"

	var aduan []models.Aduan
	if err := setup.DB.
		Preload("Program").
		Preload("User.Jabatan").
		Preload("User.Role").
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Find(&aduan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": aduan})
}

func GetAduanPerTahun(c *gin.Context) {
	tahun := c.Query("tahun")

	if tahun == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter tahun diperlukan"})
		return
	}

	startDate := tahun + "-01-01"
	endDate := tahun + "-12-31"

	var aduan []models.Aduan
	if err := setup.DB.
		Preload("Program").
		Preload("User.Jabatan").
		Preload("User.Role").
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Find(&aduan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": aduan})
}
