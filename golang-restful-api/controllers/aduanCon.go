package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"fmt"
	"net/http"
	"time"

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
			"tanggapan":  aduan.Tanggapan,
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
		Status:    "Belum Ditanggapi",
	}

	if err := tx.Create(&aduan).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat aduan"})
		return
	}

	newLog := models.Log{
		UserId:    input.UserId,
		Aktivitas: "Pengajuan Keluhan",
		Status:    "Belum Ditanggapi",
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

func TanggapiAduan(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Tanggapan string `json:"tanggapan" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tanggapan harus diisi"})
		return
	}

	var aduan models.Aduan

	if err := setup.DB.First(&aduan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aduan tidak ditemukan"})
		return
	}

	aduan.Status = "Sudah Ditanggapi"
	aduan.Tanggapan = input.Tanggapan

	tx := setup.DB.Begin()

	if err := tx.Save(&aduan).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menganggapi aduan"})
		return
	}

	newLog := models.Log{
		UserId:    aduan.UserId,
		Aktivitas: "Menganggapi Aduan",
		Status:    "Sudah Ditanggapi",
	}

	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat log aktivitas"})
		return
	}

	tx.Commit()

	if err := setup.DB.
		Preload("Program").
		Preload("User.Jabatan").
		Preload("User.Role").
		First(&aduan, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data aduan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Aduan berhasil ditanggapi",
		"data":    aduan,
	})
}

func CountAduan(c *gin.Context) {
	var count_aduan int64
	var count_aduan_belum_ditanggapi int64
	var count_aduan_sudah_ditanggapi int64

	if err := setup.DB.Model(&models.Aduan{}).Count(&count_aduan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := setup.DB.Model(&models.Aduan{}).Where("status = ?", "Belum Ditanggapi").Count(&count_aduan_belum_ditanggapi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := setup.DB.Model(&models.Aduan{}).Where("status = ?", "Sudah Ditanggapi").Count(&count_aduan_sudah_ditanggapi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": count_aduan, "belum_ditanggapi": count_aduan_belum_ditanggapi, "sudah_ditanggapi": count_aduan_sudah_ditanggapi})
}

func CountAduanPerBulan(c *gin.Context) {
	tahun := c.Query("tahun")
	if tahun == "" {
		tahun = time.Now().Format("2006") // Jika tahun tidak diisi, gunakan tahun sekarang
	}

	// Slice untuk menyimpan data bulanan
	monthlyData := make([]gin.H, 12)

	// Nama-nama bulan dalam bahasa Indonesia
	bulan := []string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}

	// Hitung untuk setiap bulan
	for i := 0; i < 12; i++ {
		bulanNum := fmt.Sprintf("%02d", i+1) // Format bulan menjadi "01", "02", dst

		// Hitung total aduan
		var totalCount int64
		setup.DB.Model(&models.Aduan{}).
			Where("YEAR(created_at) = ? AND MONTH(created_at) = ?", tahun, bulanNum).
			Count(&totalCount)

		// Simpan data bulanan
		monthlyData[i] = gin.H{
			"bulan":       bulan[i],
			"bulan_angka": bulanNum,
			"total":       totalCount,
		}
	}

	var totalKeseluruhan int64

	setup.DB.Model(&models.Aduan{}).
		Where("YEAR(created_at) = ?", tahun).
		Count(&totalKeseluruhan)

	c.JSON(http.StatusOK, gin.H{
		"tahun": tahun,
		"total_keseluruhan": gin.H{
			"total": totalKeseluruhan,
		},
		"data_bulanan": monthlyData,
	})
}

func CountAduanPerTahun(c *gin.Context) {
	// Mendapatkan tahun sekarang
	currentYear := time.Now().Year()

	// Slice untuk menyimpan data tahunan
	yearlyData := make([]gin.H, 5)

	// Loop untuk 5 tahun terakhir
	for i := 0; i < 5; i++ {
		tahun := currentYear - i

		// Format tanggal awal dan akhir tahun
		startDate := fmt.Sprintf("%d-01-01", tahun)
		endDate := fmt.Sprintf("%d-12-31", tahun)

		// Hitung total aduan
		var totalCount int64
		setup.DB.Model(&models.Aduan{}).
			Where("created_at BETWEEN ? AND ?", startDate, endDate).
			Count(&totalCount)

		// Ambil detail aduan untuk tahun ini
		var aduanList []models.Aduan
		if err := setup.DB.
			Preload("Program").
			Preload("User.Jabatan").
			Preload("User.Role").
			Where("created_at BETWEEN ? AND ?", startDate, endDate).
			Find(&aduanList).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Format data aduan
		formattedAduan := make([]gin.H, len(aduanList))
		for j, aduan := range aduanList {
			formattedAduan[j] = gin.H{
				"id":         aduan.Id,
				"program_id": aduan.ProgramId,
				"program":    aduan.Program,
				"user_id":    aduan.UserId,
				"user":       aduan.User,
				"keluhan":    aduan.Keluhan,
				"status":     aduan.Status,
				"tanggapan":  aduan.Tanggapan,
				"created_at": aduan.CreatedAt.Format("02-01-2006"),
				"updated_at": aduan.UpdatedAt.Format("02-01-2006"),
			}
		}

		// Simpan data tahunan
		yearlyData[i] = gin.H{
			"tahun":      tahun,
			"total":      totalCount,
			"data_aduan": formattedAduan,
		}
	}

	// Hitung total keseluruhan untuk 5 tahun terakhir
	var totalKeseluruhan int64

	startDate := fmt.Sprintf("%d-01-01", currentYear-4)
	endDate := fmt.Sprintf("%d-12-31", currentYear)

	setup.DB.Model(&models.Aduan{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Count(&totalKeseluruhan)

	c.JSON(http.StatusOK, gin.H{
		"periode": fmt.Sprintf("%d-%d", currentYear-4, currentYear),
		"total_keseluruhan": gin.H{
			"total": totalKeseluruhan,
		},
		"data_tahunan": yearlyData,
	})
}

func DetailAduan(c *gin.Context) {
	id := c.Param("id")

	var aduan models.Aduan

	if err := setup.DB.
		Preload("Program").
		Preload("User.Jabatan").
		Preload("User.Role").
		First(&aduan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aduan tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": aduan})
}

func GetAduanByStatus(c *gin.Context) {
	status := c.Param("status")

	validStatus := map[string]bool{
		"Belum Ditanggapi": true,
		"Sudah Ditanggapi": true,
	}

	if !validStatus[status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status tidak valid"})
		return
	}

	var aduan []models.Aduan

	if err := setup.DB.
		Preload("Program").
		Preload("User.Jabatan").
		Preload("User.Role").
		Where("status = ?", status).
		Find(&aduan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	formattedAduan := make([]gin.H, len(aduan))

	for i, aduan := range aduan {
		formattedAduan[i] = gin.H{
			"id":         aduan.Id,
			"program_id": aduan.ProgramId,
			"program":    aduan.Program,
			"user_id":    aduan.UserId,
			"user":       aduan.User,
			"keluhan":    aduan.Keluhan,
			"status":     aduan.Status,
			"tanggapan":  aduan.Tanggapan,
			"created_at": aduan.CreatedAt.Format("02-01-2006"),
			"updated_at": aduan.UpdatedAt.Format("02-01-2006"),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data":   formattedAduan,
	})
}
