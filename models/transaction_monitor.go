package models

import "time"

type TransactionMonitor struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	TransactionID   string    `json:"transaction_id" gorm:"unique;not null"`
	MSISDN          string    `json:"msisdn" gorm:"not null;index"`
	TransactionType string    `json:"transaction_type" gorm:"not null"`
	Amount          float64   `json:"amount" gorm:"type:numeric;not null"`
	Status          string    `json:"status" gorm:"type:varchar(20);not null;index"`
	CreatedAt       time.Time `json:"created_at"`
}

func (TransactionMonitor) TableName() string {
	return "transaction_monitors"
}