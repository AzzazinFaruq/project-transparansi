package models

import "time"

type Aduan struct {
	Id            int64     `gorm:"primary_key"`
	ProgramId     int64     `json:"program_id"`
	Program       Program
	UserId        int64     `json:"user_id"`
	User          User
	Keluhan       string    `json:"keluhan"`
	Status        string    `json:"status"`
	Tanggapan     string    `json:"tanggapan"`
	UserTanggapan int64     `json:"user_tanggapan"`
	TanggapanUser User      `gorm:"foreignKey:UserTanggapan;references:Id"`
	CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`                             // Menggunakan TIMESTAMP dengan default nilai CURRENT_TIMESTAMP
	UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // Menggunakan TIMESTAMP dengan auto-update
}
