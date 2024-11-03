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

// CreateJenisAnggaran menambahkan jenis anggaran baru
func CreateJenisAnggaran(c *gin.Context) {
	var input models.JenisAnggaran

	// Bind JSON ke struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	// Simpan ke database
	result := setup.DB.Create(&input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": input})
}

// DeleteJenisAnggaran menghapus jenis anggaran berdasarkan ID
func DeleteJenisAnggaran(c *gin.Context) {
	id := c.Param("id")
	
	// Cek apakah data exists
	var jenisAnggaran models.JenisAnggaran
	if err := setup.DB.First(&jenisAnggaran, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jenis anggaran tidak ditemukan"})
		return
	}
	// Hapus dari database
	if err := setup.DB.Delete(&jenisAnggaran).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Jenis anggaran berhasil dihapus"})
}
