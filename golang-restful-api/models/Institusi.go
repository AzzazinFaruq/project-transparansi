package models

import "time"

type Institusi struct {
	Id            int64     `gorm:"primary_key"`
	NamaInstitusi string    `json:"nama_institusi"`
	Alamat        string    `gorm:"type:text"`
    Email         string    `gorm:"type:varchar(255)"`
    NoTelp        string    `gorm:"type:varchar(255)"`
    Logo          string    `gorm:"type:varchar(255)"`
	CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`                             // Menggunakan TIMESTAMP dengan default nilai CURRENT_TIMESTAMP
	UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // Menggunakan TIMESTAMP dengan auto-update
}
