// this page are including all of auth controllers
package controllers

import (
	"Azzazin/backend/models"
	"Azzazin/backend/setup"
	"Azzazin/backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

	var input struct {
		Username  string `json:"username" binding:"required"`
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required,min=8"`
		NoHp      string `json:"no_hp" binding:"required"`
		Alamat    string `json:"alamat" binding:"required"`
		RoleId    int64  `json:"role_id" binding:"required"`
		JabatanId int64  `json:"jabatan_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(hashedPassword)

	user := models.User{Username: input.Username, Email: input.Email, Password: string(hashedPassword), NoHp: input.NoHp, Alamat: input.Alamat, RoleId: input.RoleId, JabatanId: input.JabatanId}

	if err := setup.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Created Succesfully"})

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
	if err := setup.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
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
	c.SetCookie("Authorization", tokenString, int(tokenDuration.Seconds()), "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message":       "Login successful",
		"username":      user.Username,
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
	c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)

	// Kirim respon logout sukses
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

