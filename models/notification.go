package models

import "time"

type Notification struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Title     string    `json:"title" gorm:"not null"`
	Message   string    `json:"message" gorm:"type:text"`
	IsRead    bool      `json:"is_read" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
}

func (Notification) TableName() string {
	return "notifications"
}