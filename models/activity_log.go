package models

import "time"

const (
	ActivityCreateBug = "CREATE_BUG"
	ActivityUpdateBug = "UPDATE_BUG"
	ActivityDeleteBug = "DELETE_BUG"
)

type ActivityLog struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	Action      string    `json:"action" gorm:"not null"`
	Description string    `json:"description"`
	IPAddress   string    `json:"ip_address"`
	CreatedAt   time.Time `json:"created_at"`
}

func (ActivityLog) TableName() string {
	return "activity_logs"
}