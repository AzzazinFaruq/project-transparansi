package models

import "time"

type Program struct {
	Id                   string `gorm:"type:char(36);primary_key;autoIncrement:false"`
	NamaProgram          string `json:"nama_program"`
	Deskripsi            string `json:"deskripsi"`
	InstitusiId          string
	Institusi            Institusi
	JenisAnggaranId      string
	JenisAnggaran        JenisAnggaran
	JumlahAnggaran       string `json:"jumlah_anggaran"`
	KategoriPenggunaanId string
	KategoriPenggunaan   KategoriPenggunaan
	FotoBefore           string
	FotoProgress         string
	FotoAfter            string
	Dusun                string
	DesaId               string
	KecamatanId          string
	KabupatenId          string
	UserId               string
	User                 User
	Status               string
	CreatedAt            time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`                             // Menggunakan TIMESTAMP dengan default nilai CURRENT_TIMESTAMP
	UpdatedAt            time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // Menggunakan TIMESTAMP dengan auto-update
}

type KategoriPenggunaan struct {
	Id       string `gorm:"type:char(36);primary_key;autoIncrement:false"`
	Kategori string `json:"kategori"`
}
type JenisAnggaran struct {
	Id    string `gorm:"type:char(36);primary_key;autoIncrement:false"`
	Jenis string `json:"jenis"`
}
