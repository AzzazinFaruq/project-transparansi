package models

import "time"

type Log struct {
	Id        int64     `gorm:"primary_key"`
	UserId    int64
	User      User
	Aktivitas string    `json:"aktivitas"`
	Status    string    `json:"status"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

