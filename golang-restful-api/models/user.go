package models

import "time"

//Structur Table User
type User struct {
	Id         int64     `gorm:"primary_key"`
	FotoProfil string    `json:"foto_profil"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	NoHp       string    `json:"no_hp"`
	Alamat     string    `json:"alamat"`
	RoleId     int64     `json:"role_id"`
	JabatanId  int64     `json:"jabatan_id"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`                             // Menggunakan TIMESTAMP dengan default nilai CURRENT_TIMESTAMP
	UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // Menggunakan TIMESTAMP dengan auto-update
	Role       Role
	Jabatan    Jabatan
}

type Role struct {
	Id   int64  `gorm:"primary_key"`
	Role string `json:"role"`
}

type Jabatan struct {
	Id      int64  `gorm:"primary_key"`
	Jabatan string `json:"jabatan"`
}
