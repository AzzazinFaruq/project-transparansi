// this page are including all of auth controllers
package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"Azzazin/backend/utils"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// Ambil data dari form
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Validasi input
	if username == "" || email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Semua field harus diisi"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
		return
	}

	// Handle foto profil jika ada
	var fotoProfil string
	file, err := c.FormFile("foto_profil")
	if err == nil {
		// Validasi tipe file
		if !strings.HasSuffix(strings.ToLower(file.Filename), ".jpg") &&
			!strings.HasSuffix(strings.ToLower(file.Filename), ".jpeg") &&
			!strings.HasSuffix(strings.ToLower(file.Filename), ".png") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung"})
			return
		}

		// Validasi ukuran file (max 5MB)
		if file.Size > 5*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ukuran file terlalu besar"})
			return
		}

		// Simpan file
		uploadPath := "uploads/profile_pictures/" + file.Filename
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan foto profil"})
			return
		}
		fotoProfil = uploadPath
	}

	// Buat user baru
	user := models.User{
		Username:   username,
		Email:      email,
		Password:   string(hashedPassword),
		FotoProfil: fotoProfil,
		RoleId:     3, // Default role masyarakat
		JabatanId:  7, // Default jabatan masyarakat
	}

	if err := setup.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat user baru"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dibuat"})
}

func Login(c *gin.Context) {
	var input struct {
		Email      string `json:"email"`
		Password   string `json:"password"`
		RememberMe bool   `json:"remember_me"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user
	var user models.User
	if err := setup.DB.Preload("Role").Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password", "authenticated": false})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password", "authenticated": false})
		return
	}

	// Tentukan durasi token berdasarkan RememberMe
	var tokenDuration time.Duration
	if input.RememberMe {
		tokenDuration = 7 * 24 * time.Hour // 7 hari jika Remember Me dicentang
	} else {
		tokenDuration = 24 * time.Hour // 1 hari
	}

	// Generate token JWT menggunakan helper
	tokenString, err := utils.GenerateJWT(uint(user.Id))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to generate token, Internal Server Error", "authenticated": false})
		return
	}

	// Set cookie dengan token JWT
	c.SetCookie("Authorization", tokenString, int(tokenDuration.Seconds()), "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message":       "Login successful",
		"username":      user.Username,
		"role":          user.Role.Role,
		"authenticated": true,
	})
}

// Fungsi untuk mendapatkan data user yang sedang login
func GetCurrentUser(c *gin.Context) {
	// Ambil user ID dari context (disimpan di middleware)
	userID, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found", "status": false})
		return
	}

	// Cari user berdasarkan ID
	var user models.User
	if err := setup.DB.Preload("Role").Preload("Jabatan").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found", "status": false})
		return
	}

	// Kembalikan data user (tanpa password untuk alasan keamanan)
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// Fungsi untuk logout
func Logout(c *gin.Context) {
	// Hapus cookie Authorization dengan durasi negatif untuk menghilangkannya
	c.SetCookie("Authorization", "", -1, "/", "", false, true)

	// Kirim respon logout sukses
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func UpdateProfile(c *gin.Context) {
	// Ambil user ID dari context (disimpan di middleware)
	userID, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak terautentikasi"})
		return
	}

	// Cari user berdasarkan ID
	var user models.User
	if err := setup.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// Ambil data dari form
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Update data user
	updates := map[string]interface{}{}

	if username != "" {
		updates["username"] = username
	}
	if email != "" {
		updates["email"] = email
	}
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
			return
		}
		updates["password"] = string(hashedPassword)
	}

	// Handle foto profil jika ada
	file, err := c.FormFile("foto_profil")
	if err == nil {
		// Validasi tipe file
		if !strings.HasSuffix(strings.ToLower(file.Filename), ".jpg") &&
			!strings.HasSuffix(strings.ToLower(file.Filename), ".jpeg") &&
			!strings.HasSuffix(strings.ToLower(file.Filename), ".png") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung"})
			return
		}

		// Validasi ukuran file (max 5MB)
		if file.Size > 5*1024*1024 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ukuran file terlalu besar"})
			return
		}

		// Hapus foto lama jika ada
		if user.FotoProfil != "" {
			os.Remove(user.FotoProfil)
		}

		// Simpan file baru
		uploadPath := "uploads/profile_pictures/" + file.Filename
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan foto profil"})
			return
		}
		updates["foto_profil"] = uploadPath
	}

	// Update data user
	if err := setup.DB.Model(&user).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui profil"})
		return
	}

	// Reload user data dengan relasi
	setup.DB.Preload("Role").Preload("Jabatan").First(&user, userID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Profil berhasil diperbarui",
		"data":    user,
	})
}
