package setup

import (
	"Azzazin/backend/models"
	"Azzazin/backend/seeders"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() { //username:password(if exist)@tcp(database URL)/database_name
	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_project_tranparansi_publik?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(models.Aduan{})
	database.AutoMigrate(models.User{}, models.Role{}, models.Institusi{}, models.Program{}, models.JenisAnggaran{}, models.KategoriPenggunaan{})

	DB = database

	seeders.SeedersUser(*DB)
}
