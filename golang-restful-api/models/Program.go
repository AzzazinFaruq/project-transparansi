package models

import "time"

type Program struct {
	Id                   int64  `gorm:"primary_key"`
	NamaProgram          string `json:"nama_program"`
	Deskripsi            string `json:"deskripsi"`
	InstitusiId          int64
	Institusi            Institusi
	JenisAnggaranId      int64
	JenisAnggaran        JenisAnggaran
	JumlahAnggaran       string `json:"jumlah_anggaran"`
	KategoriPenggunaanId int64
	KategoriPenggunaan   KategoriPenggunaan
	FotoBefore           string
	FotoProgress         string
	FotoAfter            string
	Dusun                string
	DesaId               int64
	KecamatanId          int64
	KabupatenId          int64
	UserId               int64
	User                 User
	Status               string
	CreatedAt            time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`                             // Menggunakan TIMESTAMP dengan default nilai CURRENT_TIMESTAMP
	UpdatedAt            time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // Menggunakan TIMESTAMP dengan auto-update
}

type KategoriPenggunaan struct {
	Id       int64  `gorm:"primary_key"`
	Kategori string `json:"kategori"`
}
type JenisAnggaran struct {
	Id    int64  `gorm:"primary_key"`
	Jenis string `json:"jenis"`
}
