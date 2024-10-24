package seeders

import (
	"Azzazin/backend/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedersUser(db gorm.DB) {

	var user models.User

	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	user = models.User{
		Username:  "Admin",
		Email:     "test@user.com",
		Password:  string(password),
		NoHp:      "08229485792",
		Alamat:    "Malang",
		RoleId:    1,
		JabatanId: 1,
	}

	// Cek apakah sudah ada data di tabel aduans
	if err := db.First(&user).Error; err == nil {
		log.Println("seed has been created")
		return
	}

	db.Create(&user)
	log.Println("seed success")
}
