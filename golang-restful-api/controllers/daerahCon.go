package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllKabupaten(c *gin.Context)  {
	
	var data []models.Kabupaten
	setup.DB.Find(&data)
	c.JSON(http.StatusOK, gin.H{
		
		"data": data,
	})

}

func GetAllKecamatan(c *gin.Context) {

	var data[]models.Kecamatan
	setup.DB.Preload("Kabupaten").
	Find(&data)

	c.JSON(http.StatusOK, gin.H{
		
		"data": data,
	})
}

func GetAllDesa(c *gin.Context) {

	var data[]models.Desa
	setup.DB.Preload("Kabupaten").
	Preload("Kecamatan").Find(&data)

	c.JSON(http.StatusOK, gin.H{
		
		"data": data,
	})
}
func GetKecamatanByKabupatenId(c *gin.Context) {
	var data[]models.Kecamatan
	setup.DB.Preload("Kabupaten").
	Where("id_kabupaten = ?", c.Param("id")).
	Find(&data)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func GetDesaByKecamatanId(c *gin.Context) {
	var data[]models.Desa
	setup.DB.Preload("Kecamatan").
	Preload("Kabupaten").
	Where("id_kecamatan = ?", c.Param("id")).
	Find(&data)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
