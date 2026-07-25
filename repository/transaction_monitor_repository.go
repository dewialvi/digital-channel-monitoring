package repository

import (
	"gorm.io/gorm"

	"github.com/dewialvi/digital-channel-monitoring/models"
)

type TransactionMonitorRepository struct {
	DB *gorm.DB
}

func NewTransactionMonitorRepository(
	db *gorm.DB,
) *TransactionMonitorRepository {
	return &TransactionMonitorRepository{
		DB: db,
	}
}

func (r *TransactionMonitorRepository) Create(
	t *models.TransactionMonitor,
) error {
	return r.DB.Create(t).Error
}

type TransactionFilter struct {
	MSISDN string
	Status string
	Page   int
	Limit  int
}

func (r *TransactionMonitorRepository) FindAll(
	filter TransactionFilter,
) ([]models.TransactionMonitor, int64, error) {

	var trx []models.TransactionMonitor
	var total int64

	query := r.DB.Model(&models.TransactionMonitor{})

	if filter.MSISDN != "" {
		query = query.Where(
			"msisdn = ?",
			filter.MSISDN,
		)
	}

	if filter.Status != "" {
		query = query.Where(
			"status = ?",
			filter.Status,
		)
	}

	query.Count(&total)

	offset := (filter.Page - 1) * filter.Limit

	err := query.
		Order("created_at DESC").
		Limit(filter.Limit).
		Offset(offset).
		Find(&trx).
		Error

	return trx, total, err
}