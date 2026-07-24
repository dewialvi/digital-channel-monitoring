package models

import "time"

type APIMonitor struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Endpoint       string    `json:"endpoint" gorm:"not null;index"`
	Method         string    `json:"method" gorm:"type:varchar(10);not null"`
	StatusCode     int       `json:"status_code" gorm:"not null"`
	ResponseTimeMs int       `json:"response_time_ms" gorm:"not null"`
	CheckedAt      time.Time `json:"checked_at" gorm:"not null;index"`
	CreatedAt      time.Time `json:"created_at"`
}

func (APIMonitor) TableName() string {
	return "api_monitors"
}