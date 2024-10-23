package models

import "time"

type Aduan struct {
	Id        string `gorm:"type:char(36);primary_key;autoIncrement:false"`
	ProgramId string
	Program   Program
	UserId    string
	User      User
	Keluhan   string    `json:"keluhan"`
	Status    string    `json:"status"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`                             // Menggunakan TIMESTAMP dengan default nilai CURRENT_TIMESTAMP
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // Menggunakan TIMESTAMP dengan auto-update
}
