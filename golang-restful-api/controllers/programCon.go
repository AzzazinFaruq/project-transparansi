package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProgram(c *gin.Context) {
	var program []models.Program

	if err := setup.DB.
		Preload("Institusi").
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("User.Jabatan").
		Preload("User.Role").
		Find(&program).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	formattedPrograms := make([]gin.H, len(program))

	for i, program := range program {
		formattedPrograms[i] = gin.H{
			"id":                  program.Id,
			"nama_program":        program.NamaProgram,
			"institusi":           program.Institusi,
			"jenis_anggaran":      program.JenisAnggaran,
			"kategori_penggunaan": program.KategoriPenggunaan,
			"user":                program.User,
			"status":              program.Status,
			"created_at":          program.CreatedAt.Format("02-01-2006"),
			"updated_at":          program.UpdatedAt.Format("02-01-2006"),
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": formattedPrograms})
}

func PengajuanProgram(c *gin.Context) {
	var input models.Program

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid"})
		return
	}

	newProgram := models.Program{
		NamaProgram:          input.NamaProgram,
		Deskripsi:            input.Deskripsi,
		InstitusiId:          input.InstitusiId,
		JenisAnggaranId:      input.JenisAnggaranId,
		JumlahAnggaran:       input.JumlahAnggaran,
		KategoriPenggunaanId: input.KategoriPenggunaanId,
		Dusun:                input.Dusun,
		DesaId:               input.DesaId,
		KecamatanId:          input.KecamatanId,
		KabupatenId:          input.KabupatenId,
		UserId:               input.UserId,
		Status:               "Menunggu",
	}

	tx := setup.DB.Begin()

	if err := tx.Create(&newProgram).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengajukan program baru"})
		return
	}

	newLog := models.Log{
		UserId:    input.UserId,
		Aktivitas: "Pengajuan Program",
		Status:    "Menunggu",
	}

	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat log aktivitas"})
		return
	}

	tx.Commit()

	setup.DB.Preload("Institusi").
		Preload("JenisAnggaran").
		Preload("KategoriPenggunaan").
		Preload("User").
		First(&newProgram, newProgram.Id)

	c.JSON(http.StatusCreated, gin.H{"message": "Program berhasil diajukan", "data": newProgram})
}

func EditProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program
	var input models.Program

	if err := setup.DB.First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid"})
		return
	}

	tx := setup.DB.Begin()

	updateData := models.Program{
		NamaProgram:          input.NamaProgram,
		Deskripsi:            input.Deskripsi,
		InstitusiId:          input.InstitusiId,
		JenisAnggaranId:      input.JenisAnggaranId,
		JumlahAnggaran:       input.JumlahAnggaran,
		KategoriPenggunaanId: input.KategoriPenggunaanId,
		FotoBefore:           input.FotoBefore,
		Dusun:                input.Dusun,
		DesaId:               input.DesaId,
		KecamatanId:          input.KecamatanId,
		KabupatenId:          input.KabupatenId,
	}

	if err := tx.Model(&program).Updates(updateData).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate program"})
		return
	}

	newLog := models.Log{
		UserId:    program.UserId,
		Aktivitas: "Edit Program",
		Status:    program.Status,
	}

	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat log aktivitas"})
		return
	}

	tx.Commit()

	setup.DB.Preload("Institusi").
		Preload("JenisAnggaran").
		Preload("KategoriPenggunaan").
		Preload("User").
		First(&program, program.Id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Program berhasil diupdate",
		"data":    program,
	})
}

func AcceptProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if err := setup.DB.First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program tidak ditemukan"})
		return
	}

	program.Status = "Dalam Proses"

	tx := setup.DB.Begin()

	if err := tx.Save(&program).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyetujui program"})
		return
	}

	newLog := models.Log{
		UserId:    program.UserId,
		Aktivitas: "Penyetujuan Program",
		Status:    "Dalam Proses",
	}

	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat log aktivitas"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Program berhasil disetujui", "data": program})
}

func RejectProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if err := setup.DB.Find(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program tidak ditemukan"})
		return
	}

	program.Status = "Ditolak"

	tx := setup.DB.Begin()

	if err := tx.Save(&program).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menolak program"})
		return
	}

	newLog := models.Log{
		UserId:    program.UserId,
		Aktivitas: "Penolakan Program",
		Status:    "Ditolak",
	}

	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat log aktivitas"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Program berhasil ditolak", "data": program})
}

func GetProgramByStatus(c *gin.Context) {
	status := c.Param("status")

	validStatus := map[string]bool{
		"Menunggu":  true,
		"Disetujui": true,
		"Ditolak":   true,
	}

	if !validStatus[status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status tidak valid"})
		return
	}

	var programs []models.Program

	if err := setup.DB.
		Preload("Institusi").
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("User.Jabatan").
		Preload("User.Role").
		Where("status = ?", status).
		Find(&programs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	formattedPrograms := make([]gin.H, len(programs))

	for i, program := range programs {
		formattedPrograms[i] = gin.H{
			"id":                  program.Id,
			"nama_program":        program.NamaProgram,
			"institusi":           program.Institusi,
			"jenis_anggaran":      program.JenisAnggaran,
			"jumlah_anggaran":     program.JumlahAnggaran,
			"kategori_penggunaan": program.KategoriPenggunaan,
			"user":                program.User,
			"status":              program.Status,
			"created_at":          program.CreatedAt.Format("02-01-2006"),
			"updated_at":          program.UpdatedAt.Format("02-01-2006"),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data":   formattedPrograms,
	})
}

func DetailProgram(c *gin.Context) {
	id := c.Param("id")

	var program models.Program

	if err := setup.DB.
		Preload("Institusi").
		Preload("KategoriPenggunaan").
		Preload("JenisAnggaran").
		Preload("User.Jabatan").
		Preload("User.Role").
		First(&program, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": program})
}
