package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)



func GetInstitusiByNamaInstitusi(c *gin.Context) {
	// Ambil parameter pencarian dari query string
	query := c.Query("nama_institusi")

	// Jika query kosong
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pencarian tidak boleh kosong"})
		return
	}

	var institusi []models.Institusi

	// Gunakan LIKE untuk pencarian partial match
	if err := setup.DB.
		Where("nama_institusi LIKE ?", "%"+query+"%").
		Find(&institusi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Jika tidak ada hasil
	if len(institusi) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Tidak ada institusi yang ditemukan",
			"data":    []interface{}{},
		})
		return
	}

	// Format hasil pencarian
	formattedInstitusi := make([]gin.H, len(institusi))
	for i, institusi := range institusi {
		formattedInstitusi[i] = gin.H{
			"id":                  institusi.Id,
			"nama_institusi":      institusi.NamaInstitusi,
			"created_at":          institusi.CreatedAt.Format("02-01-2006"),
			"updated_at":          institusi.UpdatedAt.Format("02-01-2006"),	
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Institusi ditemukan",
		"total":   len(institusi),
		"data":    formattedInstitusi,
	})
}

func CreateInstitusi(c *gin.Context) {
	// Ambil file
	file, err := c.FormFile("logo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Logo tidak ditemukan"})
		return
	}

	// Generate nama file unik
	filename := filepath.Base(file.Filename)
	finalPath := "uploads/institusi/" + filename

	// Simpan file
	if err := c.SaveUploadedFile(file, "public/"+finalPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan logo"})
		return
	}

	// Buat institusi baru
	institusi := models.Institusi{
		NamaInstitusi: c.PostForm("nama_institusi"),
		Alamat:        c.PostForm("alamat"),
		Email:         c.PostForm("email"),
		NoTelp:        c.PostForm("no_telp"),
		Logo:          finalPath,
	}

	if err := setup.DB.Create(&institusi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": institusi})
}

func GetInstitusi(c *gin.Context) {
	var institusi []models.Institusi

	if err := setup.DB.Find(&institusi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	formattedInstitusi := make([]gin.H, len(institusi))

	for i, institusi := range institusi {
		formattedInstitusi[i] = gin.H{
			"id": institusi.Id,
			"nama_institusi": institusi.NamaInstitusi,
			"created_at": institusi.CreatedAt.Format("02-01-2006"),
			"updated_at": institusi.UpdatedAt.Format("02-01-2006"),
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": formattedInstitusi})
}

func GetInstitusiById(c *gin.Context) {
	id := c.Param("id")
	var institusi models.Institusi

	if err := setup.DB.First(&institusi, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Institusi tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": institusi})
}

func UpdateInstitusi(c *gin.Context) {
	id := c.Param("id")
	var institusi models.Institusi

	// Cari institusi yang akan diupdate
	if err := setup.DB.First(&institusi, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Institusi tidak ditemukan"})
		return
	}

	// Update data
	institusi.NamaInstitusi = c.PostForm("nama_institusi")
	institusi.Alamat = c.PostForm("alamat")
	institusi.Email = c.PostForm("email")
	institusi.NoTelp = c.PostForm("no_telp")

	// Jika ada file logo baru
	if file, err := c.FormFile("logo"); err == nil {
		filename := filepath.Base(file.Filename)
		finalPath := "uploads/institusi/" + filename

		if err := c.SaveUploadedFile(file, "public/"+finalPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan logo"})
			return
		}

		institusi.Logo = finalPath
	}

	if err := setup.DB.Save(&institusi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": institusi})
}

func DeleteInstitusi(c *gin.Context) {
	id := c.Param("id")
	var institusi models.Institusi

	if err := setup.DB.First(&institusi, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Institusi tidak ditemukan"})
		return
	}

	if err := setup.DB.Delete(&institusi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Institusi berhasil dihapus"})
}