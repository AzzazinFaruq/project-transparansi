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
		Preload("TanggapanUser").
		Order("created_at DESC").
		Find(&Aduan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	formattedAduans := make([]gin.H, len(Aduan))

	for i, aduan := range Aduan {
		formattedAduans[i] = gin.H{
			"id":             aduan.Id,
			"program_id":     aduan.ProgramId,
			"program":        aduan.Program,
			"user_id":        aduan.UserId,
			"user":           aduan.User,
			"keluhan":        aduan.Keluhan,
			"status":         aduan.Status,
			"tanggapan":      aduan.Tanggapan,
			"user_tanggapan": aduan.UserTanggapan,
			"created_at":     aduan.CreatedAt.Format("02-01-2006"), // Format d-m-y
			"updated_at":     aduan.UpdatedAt.Format("02-01-2006"), // Format d-m-y
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": formattedAduans})
}

func CreateAduan(c *gin.Context) {
	var input struct {
		ProgramId int64  `json:"program_id" binding:"required"`
		UserId    int64  `json:"user_id" binding:"required"`
		Keluhan   string `json:"keluhan" binding:"required"`
		NoHp      string `json:"no_hp" binding:"required"`
		Alamat    string `json:"alamat" binding:"required"`
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
		NoHp:      input.NoHp,
		Alamat:    input.Alamat,
		Status:    "Belum Ditanggapi",
		Tanggapan: nil,
		UserTanggapan: nil,
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
		Tanggapan     string `json:"tanggapan" binding:"required"`
		UserTanggapan int64  `json:"user_tanggapan" binding:"required"`
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
	aduan.Tanggapan = &input.Tanggapan
	aduan.UserTanggapan = &input.UserTanggapan
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
		tahun = time.Now().Format("2006") // Default tahun sekarang
	}

	// Slice untuk menyimpan jumlah aduan per bulan
	var monthlyCount []int64

	// Nama-nama bulan
	bulan := []string{
		"Jan", "Feb", "Mar", "Apr", "Mei", "Jun",
		"Jul", "Agu", "Sep", "Okt", "Nov", "Des",
	}

	// Hitung aduan untuk setiap bulan
	for i := 1; i <= 12; i++ {
		bulanNum := fmt.Sprintf("%02d", i)
		var count int64

		setup.DB.Model(&models.Aduan{}).
			Where("YEAR(created_at) = ? AND MONTH(created_at) = ?", tahun, bulanNum).
			Count(&count)

		monthlyCount = append(monthlyCount, count)
	}

	// Format response untuk ApexChart
	chartData := gin.H{
		"series": []gin.H{
			{
				"name": "Jumlah Aduan",
				"data": monthlyCount,
			},
		},
		"options": gin.H{
			"chart": gin.H{
				"height": 350,
				"type":   "line",
				"zoom": gin.H{
					"enabled": false,
				},
			},
			"stroke": gin.H{
				"curve": "straight",
			},
			"title": gin.H{
				"align": "left",
			},
			"grid": gin.H{
				"row": gin.H{
					"colors":  []string{"#FFE3E3", "transparent"},
					"opacity": 0.5,
				},
			},
			"xaxis": gin.H{
				"categories": bulan,
			},
			"yaxis": gin.H{
				"min":        0,
				"max":        getMaxValue(monthlyCount) + 5,
				"tickAmount": 4,
			},
		},
	}

	c.JSON(http.StatusOK, chartData)
}

// Fungsi helper untuk mendapatkan nilai maksimum
func getMaxValue(numbers []int64) int64 {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func CountAduanPerTahun(c *gin.Context) {
	// Mendapatkan tahun sekarang
	currentYear := time.Now().Year()

	// Slice untuk menyimpan data tahunan
	var yearlyCount []int64
	var years []string

	// Loop untuk 5 tahun terakhir
	for i := 4; i >= 0; i-- { // Dibalik agar urutan dari tahun terlama ke terbaru
		tahun := currentYear - i
		years = append(years, fmt.Sprintf("%d", tahun))

		// Format tanggal awal dan akhir tahun
		startDate := fmt.Sprintf("%d-01-01", tahun)
		endDate := fmt.Sprintf("%d-12-31", tahun)

		// Hitung total aduan
		var totalCount int64
		setup.DB.Model(&models.Aduan{}).
			Where("created_at BETWEEN ? AND ?", startDate, endDate).
			Count(&totalCount)

		yearlyCount = append(yearlyCount, totalCount)
	}

	// Format response untuk ApexChart
	chartData := gin.H{
		"series": []gin.H{
			{
				"name": "Jumlah Aduan",
				"data": yearlyCount,
			},
		},
		"options": gin.H{
			"chart": gin.H{
				"height": 350,
				"type":   "line",
				"zoom": gin.H{
					"enabled": false,
				},
			},
			"stroke": gin.H{
				"curve": "straight",
			},
			"grid": gin.H{
				"row": gin.H{
					"colors":  []string{"#FFE3E3", "transparent"},
					"opacity": 0.5,
				},
			},
			"xaxis": gin.H{
				"categories": years,
			},
			"yaxis": gin.H{
				"min":        0,
				"max":        getMaxValue(yearlyCount) + 5,
				"tickAmount": 4,
			},
		},
	}

	c.JSON(http.StatusOK, chartData)
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
			"no_hp":      aduan.NoHp,
			"alamat":     aduan.Alamat,
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

func GetAduanByProgramId(c *gin.Context) {
	programId := c.Param("program_id")

	var aduan []models.Aduan

	if err := setup.DB.
		Preload("Program").
		Preload("User.Jabatan").
		Preload("User.Role").
		Preload("TanggapanUser").
		Where("program_id = ?", programId).
		Order("created_at DESC").
		Find(&aduan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	formattedAduan := make([]gin.H, len(aduan))

	for i, aduan := range aduan {
		var userTanggapan models.User
		if aduan.UserTanggapan != nil && *aduan.UserTanggapan != 0 {
			setup.DB.
				Preload("Jabatan").
				Preload("Role").
				First(&userTanggapan, aduan.UserTanggapan)
		}

		formattedAduan[i] = gin.H{
			"id":             aduan.Id,
			"program_id":     aduan.ProgramId,
			"program":        aduan.Program,
			"user_id":        aduan.UserId,
			"user":           aduan.User,
			"no_hp":          aduan.NoHp,
			"alamat":         aduan.Alamat,
			"keluhan":        aduan.Keluhan,
			"status":         aduan.Status,
			"tanggapan":      aduan.Tanggapan,
			"user_tanggapan": userTanggapan,
			"created_at":     aduan.CreatedAt.Format("02-01-2006"),
			"updated_at":     aduan.UpdatedAt.Format("02-01-2006"),
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": formattedAduan})
}
