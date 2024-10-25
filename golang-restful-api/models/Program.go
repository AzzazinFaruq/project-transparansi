package models

import "time"

type Program struct {
	Id                   int64  `gorm:"primary_key"`
	NamaProgram          string `json:"nama_program"`
	Deskripsi            string `json:"deskripsi"`
	InstitusiId          int64  `json:"institusi_id"`
	Institusi            Institusi
	JenisAnggaranId      int64  `json:"jenis_anggaran_id"`
	JenisAnggaran        JenisAnggaran
	JumlahAnggaran       string `json:"jumlah_anggaran"`
	KategoriPenggunaanId int64  `json:"kategori_penggunaan_id"`
	KategoriPenggunaan   KategoriPenggunaan
	FotoBefore           string `json:"foto_before"`
	FotoProgress         string `json:"foto_progress"`
	FotoAfter            string `json:"foto_after"`
	Dusun                string `json:"dusun"`
	DesaId               int64  `json:"desa_id"`
	KecamatanId          int64  `json:"kecamatan_id"`
	KabupatenId          int64  `json:"kabupaten_id"`
	UserId               int64  `json:"user_id"`
	User                 User
	Status               string `json:"status"`
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
