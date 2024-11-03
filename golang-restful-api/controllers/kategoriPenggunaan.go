package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllKategoriPenggunaan(c *gin.Context) {
	var kategoriPenggunaan []models.KategoriPenggunaan

	if err := setup.DB.
		Order("nama ASC").
		Find(&kategoriPenggunaan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": kategoriPenggunaan})
}

// CreateKategoriPenggunaan menambahkan kategori penggunaan baru
func CreateKategoriPenggunaan(c *gin.Context) {
	var input models.KategoriPenggunaan

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

// DeleteKategoriPenggunaan menghapus kategori penggunaan berdasarkan ID
func DeleteKategoriPenggunaan(c *gin.Context) {
	id := c.Param("id")
	
	// Cek apakah data exists
	var kategori models.KategoriPenggunaan
	if err := setup.DB.First(&kategori, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori penggunaan tidak ditemukan"})
		return
	}

	// Hapus dari database
	if err := setup.DB.Delete(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kategori penggunaan berhasil dihapus"})
}
