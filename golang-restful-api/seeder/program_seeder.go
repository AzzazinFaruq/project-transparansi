package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/bxcodec/faker/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Program struct {
	Id                    int64   `gorm:"primaryKey"`
	NamaProgram          string
	NamaInstitusi        string
	JenisAnggaranId      int64
	JumlahAnggaran       string
	KategoriPenggunaanId int64
	AspiratorId          int64
	DinasVerifikatorId   int64
	Deskripsi            string
	Dusun                string
	DesaId               int64
	KecamatanId          int64
	KabupatenId          int64
	FotoBefore           string
	FotoProgress         string
	FotoAfter            string
	UserId               int64
	Status               string
	Latitude             float64
	Longitude            float64
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func generateRandomName() string {
	names := []string{
		"Program Pembangunan Jalan",
		"Program Renovasi Sekolah",
		"Program Perbaikan Irigasi",
		"Program Pembangunan Masjid",
		"Program Renovasi Posyandu",
		"Program Pembangunan Jembatan",
		"Program Perbaikan Drainase",
		"Program Pembangunan Taman",
		"Program Renovasi Balai Desa",
		"Program Pembangunan Sanitasi",
	}
	return names[rand.Intn(len(names))] + " " + faker.Word()
}

func generateRandomInstitusi() string {
	institusi := []string{
		"Dinas Pekerjaan Umum",
		"Dinas Pendidikan",
		"Dinas Pertanian",
		"Dinas Kesehatan",
		"Dinas Sosial",
		"Dinas Lingkungan Hidup",
		"Dinas Perhubungan",
		"Dinas Perumahan",
		"Dinas Pemberdayaan Masyarakat",
		"Dinas Koperasi dan UMKM",
	}
	return institusi[rand.Intn(len(institusi))]
}

func main() {
	// Koneksi database
	dsn := "root:@tcp(127.0.0.1:3306)/db_project_tranparansi_publik?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	// Set random seed
	rand.Seed(time.Now().UnixNano())

	// Data untuk random - sesuaikan dengan ID yang ada di database
	status := []string{"Draft", "Publish"}
	jenisAnggaranIds := []int64{3,4,5} // sesuaikan dengan data di tabel jenis_anggarans
	kategoriPenggunaanIds := []int64{3,4,5} // sesuaikan dengan data di tabel kategori_penggunaans
	aspiratorIds := []int64{125,126,127,128,129,130} // sesuaikan dengan data di tabel aspirators
	dinasVerifikatorIds := []int64{4,5,6,7,8,9,10,11,12,13,14,15} // sesuaikan dengan data di tabel dinas_verifikators
	desaIds := []int64{3501010001, 3501010002, 3501010003} // sesuaikan dengan data di tabel desas
	kecamatanIds := []int64{3501010} // sesuaikan dengan data di tabel kecamatans
	kabupatenIds := []int64{3501} // sesuaikan dengan data di tabel kabupatens
	userIds := []int64{16} // sesuaikan dengan data di tabel users

	// Generate 100 program
	for i := 0; i < 100; i++ {
		randomTime := time.Now().AddDate(0, 0, -rand.Intn(365))
		
		program := Program{
			NamaProgram:          generateRandomName(),
			NamaInstitusi:        generateRandomInstitusi(),
			JenisAnggaranId:      jenisAnggaranIds[rand.Intn(len(jenisAnggaranIds))],
			JumlahAnggaran:       fmt.Sprintf("%d", rand.Int63n(990000000)+10000000),
			KategoriPenggunaanId: kategoriPenggunaanIds[rand.Intn(len(kategoriPenggunaanIds))],
			AspiratorId:          aspiratorIds[rand.Intn(len(aspiratorIds))],
			DinasVerifikatorId:   dinasVerifikatorIds[rand.Intn(len(dinasVerifikatorIds))],
			Deskripsi:            faker.Paragraph(),
			Dusun:                faker.Word(),
			DesaId:               desaIds[rand.Intn(len(desaIds))],
			KecamatanId:          kecamatanIds[rand.Intn(len(kecamatanIds))],
			KabupatenId:          kabupatenIds[rand.Intn(len(kabupatenIds))],
			UserId:               userIds[rand.Intn(len(userIds))],
			Status:               status[rand.Intn(len(status))],
			Latitude:             -7.5 + rand.Float64()*0.3,
			Longitude:            112.5 + rand.Float64()*0.3,
			CreatedAt:            randomTime,
			UpdatedAt:            randomTime,
		}

		if err := db.Create(&program).Error; err != nil {
			fmt.Printf("Error creating program: %v\n", err)
		}
	}

	fmt.Println("Successfully seeded 100 programs")
} 