package seeders

import (
	"Azzazin/backend/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedersUser(db gorm.DB) {

	var user models.User
	var role models.Role
	var jabatan models.Jabatan

	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	role = models.Role{
		Role: "admin",
	}
	jabatan = models.Jabatan{
		Jabatan: "Ketua DPRD",
	}

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
	if err := db.First(&role).Error; err == nil {
		log.Println("seed has been created")
		return
	}
	if err := db.First(&jabatan).Error; err == nil {
		log.Println("seed has been created")
		return
	}

	db.Create(&role)
	db.Create(&jabatan)
	db.Create(&user)
	log.Println("seed success")
}
