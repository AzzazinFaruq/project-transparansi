package models

import "time"

//Structur Table User
type User struct {
	Id        string `gorm:"type:char(36);primary_key;autoIncrement:false"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	NoHp      string `json:"no_hp"`
	Alamat    string `json:"alamat"`
	RoleId    string
	JabatanId string
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`                             // Menggunakan TIMESTAMP dengan default nilai CURRENT_TIMESTAMP
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // Menggunakan TIMESTAMP dengan auto-update
	Role      Role
	Jabatan   Jabatan
}

type Role struct {
	Id   string `gorm:"type:char(36);primary_key;autoIncrement:false"`
	Role string `json:"role"`
}

type Jabatan struct {
	Id      string `gorm:"type:char(36);primary_key;autoIncrement:false"`
	Jabatan string `json:"jabatan"`
}
