package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"Azzazin/backend/utils"
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

func GetDetailUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := setup.DB.
		Preload("Role").
		Preload("Jabatan").
		First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUserByRole(c *gin.Context) {
	roleId := c.Param("id")

	var users []models.User

	if err := setup.DB.
		Preload("Role").
		Preload("Jabatan").
		Joins("JOIN roles ON users.role_id = roles.id").
		Where("role_id = ?", roleId).
		Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data pengguna"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tidak ada user dengan role ini"})
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

func CreateDPRD(c *gin.Context) {
	// Mengambil data dari form yang dikirim oleh front-end
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	noHp := c.PostForm("no_hp")
	alamat := c.PostForm("alamat")
	jabatanId := c.PostForm("jabatan_id")

	if username == "" || email == "" || password == "" || noHp == "" || alamat == "" || jabatanId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Semua field harus diisi"})
		return
	}

	// Mengambil file foto profil dari form
	file, err := c.FormFile("foto_profil")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal mengambil file foto profil"})
		return
	}

	// Mengunggah file foto profil
	uploadPath := "uploads/profile_pictures"
	filePath, err := utils.UploadFile(c, file, uploadPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupload foto profil"})
		return
	}

	// Mengkonversi jabatanId dari string ke int64
	jabatanIdInt, err := strconv.ParseInt(jabatanId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format jabatan_id tidak valid"})
		return
	}

	// Memeriksa apakah jabatan yang dipilih ada di database
	var jabatan models.Jabatan
	if err := setup.DB.First(&jabatan, jabatanIdInt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Jabatan tidak ditemukan"})
		return
	}

	// Mencari role DPRD di database
	var dprdRole models.Role
	if err := setup.DB.Where("role = ?", "DPRD").First(&dprdRole).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Role DPRD tidak ditemukan"})
		return
	}

	// Mengenkripsi password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
		return
	}

	// Membuat user baru dengan data yang diterima
	newUser := models.User{
		Username:   username,
		Email:      email,
		Password:   string(hashedPassword),
		NoHp:       noHp,
		Alamat:     alamat,
		RoleId:     dprdRole.Id,
		JabatanId:  jabatanIdInt,
		FotoProfil: filePath,
	}

	// Menyimpan user baru ke database
	if err := setup.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan DPRD baru: " + err.Error()})
		return
	}

	// Mengambil data user yang baru dibuat beserta relasi Role dan Jabatan
	if err := setup.DB.Preload("Role").Preload("Jabatan").First(&newUser, newUser.Id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data DPRD yang baru dibuat"})
		return
	}

	// Mengirim respons sukses beserta data user yang baru dibuat
	c.JSON(http.StatusCreated, gin.H{"message": "DPRD berhasil dibuat", "data": newUser})
}
