package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllInstitusi(c *gin.Context) {
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

func CreateInstitusi(c *gin.Context) {
	var input struct {
		NamaInstitusi string `json:"nama_institusi" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	

	institusi := models.Institusi{
		NamaInstitusi: input.NamaInstitusi,
	}

	if err := setup.DB.Create(&institusi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

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