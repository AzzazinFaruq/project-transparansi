package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	var user []models.User

	if err := setup.DB.
		Preload("Role").
		Preload("Jabatan").
		Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUserByRole(c *gin.Context) {
	roleId := c.Param("roleId")

	var users []models.User

	if err := setup.DB.
		Preload("Role").
		Preload("Jabatan").
		Joins("JOIN roles ON users.role_id = roles.id").
		Where("roles.id = ?", roleId).
		Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data pengguna"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tidak ada pengguna dengan peran tersebut"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}
func AddUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data input tidak valid"})
		return
	}

	newUser := models.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  input.Password,
		NoHp:      input.NoHp,
		Alamat:    input.Alamat,
		RoleId:    input.RoleId,
		JabatanId: input.JabatanId,
	}

	if err := setup.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan pengguna baru"})
		return
	}

	setup.DB.Preload("Role").Preload("Jabatan").First(&newUser, newUser.Id)

	c.JSON(http.StatusCreated, gin.H{"data": newUser})
}
func EditUser(c *gin.Context) {
	userId := c.Param("id")
	var user models.User
	var input models.User

	if err := setup.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := setup.DB.Model(&user).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui pengguna"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	var user models.User

	if err := setup.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	if err := setup.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus pengguna"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengguna berhasil dihapus"})
}