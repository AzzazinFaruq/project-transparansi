package models

import "time"

type Program struct {
	Id                     int64              `gorm:"primary_key"`
	NamaProgram            string             `json:"nama_program"`
	Deskripsi              string             `json:"deskripsi"`
	NamaInstitusi          string             `json:"nama_institusi"`
	JenisAnggaranId        int64             `json:"jenis_anggaran_id"`
	JenisAnggaran          JenisAnggaran      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	JumlahAnggaran         string             `json:"jumlah_anggaran"`
	KategoriPenggunaanId   int64             `json:"kategori_penggunaan_id"`
	KategoriPenggunaan     KategoriPenggunaan `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AspiratorId            int64              `json:"aspirator_id"`
	Aspirator              Aspirator          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DinasVerifikatorId     int64              `json:"dinas_verifikator_id"`
	DinasVerifikator       DinasVerifikator   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FotoBefore             string             `json:"foto_before"`
	FotoProgress           string             `json:"foto_progress"`
	FotoAfter              string             `json:"foto_after"`
	Dusun                  string             `json:"dusun"`
	DesaId                 int64              `json:"desa_id"`
	Desa                   Desa               `gorm:"foreignKey:DesaId;references:Id"`
	KecamatanId            int64              `json:"kecamatan_id"`
	Kecamatan              Kecamatan          `gorm:"foreignKey:KecamatanId;references:Id"`
	KabupatenId            int64              `json:"kabupaten_id"`
	Kabupaten              Kabupaten          `gorm:"foreignKey:KabupatenId;references:Id"`
	UserId                 int64              `json:"user_id"`
	User                   User               `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status                 string             `json:"status"`
	CreatedAt              time.Time          `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`                             // Menggunakan TIMESTAMP dengan default nilai CURRENT_TIMESTAMP
	UpdatedAt              time.Time          `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // Menggunakan TIMESTAMP dengan auto-update
	JenisAnggaranLain      string             `json:"jenis_anggaran_lain"`
	KategoriPenggunaanLain string             `json:"kategori_penggunaan_lain"`
}

type KategoriPenggunaan struct {
	Id       int64  `gorm:"primary_key"`
	Kategori string `json:"kategori"`
}
type JenisAnggaran struct {
	Id    int64  `gorm:"primary_key"`
	Jenis string `json:"jenis"`
}

type Aspirator struct {
	Id            int64  `gorm:"primary_key"`
	NamaAspirator string `json:"nama_aspirator"`
}

type DinasVerifikator struct {
	Id                   int64  `gorm:"primary_key"`
	NamaDinasVerifikator string `json:"nama_dinas_verifikator"`
}
