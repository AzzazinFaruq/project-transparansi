package models

import "time"

type Log struct {
	Id        int64     `gorm:"primary_key"`
	UserId    int64     `json:"user_id"`
	User      User
	Aktivitas string    `json:"aktivitas"`
	Status    string    `json:"status"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

// Tambahkan metode baru untuk memformat tanggal
func (l *Log) FormattedCreatedAt() string {
	return l.CreatedAt.Format("02-01-2006") // Format d-m-y
}
