package controllers

import (
	"Azzazin/backend/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Kabupaten(c *gin.Context)  {
	
	var data []models.Kabupaten
	models.DB.Find(&data)
	c.JSON(http.StatusOK, gin.H{
		
		"data": data,
	})

}

func Kecamatan(c *gin.Context) {

	var data[]models.Kecamatan
	models.DB.Preload("Kabupaten").Find(&data)

	c.JSON(http.StatusOK, gin.H{
		
		"data": data,
	})
}

func Desa(c *gin.Context) {

	var data[]models.Desa
	models.DB.Preload("Kabupaten").Preload("Kecamatan").Find(&data)

	c.JSON(http.StatusOK, gin.H{
		
		"data": data,
	})
}