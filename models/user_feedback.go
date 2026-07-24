package models

import "time"

type UserFeedback struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Category  string    `json:"category" gorm:"not null"`
	Message   string    `json:"message" gorm:"type:text;not null"`
	Rating    int       `json:"rating" gorm:"check:rating >= 1 AND rating <= 5"`
	CreatedAt time.Time `json:"created_at"`
}

func (UserFeedback) TableName() string {
	return "user_feedbacks"
}