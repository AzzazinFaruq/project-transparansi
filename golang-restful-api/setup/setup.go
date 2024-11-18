package setup

import (
	"log"
	"os"

	"Azzazin/backend/models"
	"Azzazin/backend/seeders"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using environment variables.")
	}

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("Database configuration variables are not set properly.")
	}
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName +
		"?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Automigrate your models
	err = database.AutoMigrate(
		&models.Aduan{},
		&models.User{},
		&models.Role{},
		&models.Program{},
		&models.JenisAnggaran{},
		&models.KategoriPenggunaan{},
		&models.Log{},
		&models.Kabupaten{},
		&models.Kecamatan{},
		&models.Desa{},
	)
	if err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}

	DB = database

	// Seed the database
	seeders.SeedersUser(*DB)
}
