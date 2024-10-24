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

	c.JSON(http.StatusOK, gin.H{"data": program})
}

func PengajuanProgram(c *gin.Context) {
	var input models.Program

	// Mengikat input JSON ke struktur Program
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
		FotoBefore:           input.FotoBefore,
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
