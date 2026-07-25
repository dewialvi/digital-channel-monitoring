package repository

import (
	"gorm.io/gorm"

	"github.com/dewialvi/digital-channel-monitoring/models"
)

type ActivityLogRepository struct {
	DB *gorm.DB
}

func NewActivityLogRepository(db *gorm.DB) *ActivityLogRepository {
	return &ActivityLogRepository{
		DB: db,
	}
}

func (r *ActivityLogRepository) Create(log *models.ActivityLog) error {
	return r.DB.Create(log).Error
}

type ActivityLogFilter struct {
	Page  int
	Limit int
}

func (r *ActivityLogRepository) FindAll(
	filter ActivityLogFilter,
) ([]models.ActivityLog, int64, error) {

	var logs []models.ActivityLog
	var total int64

	query := r.DB.
		Model(&models.ActivityLog{}).
		Preload("User")

	query.Count(&total)

	offset := (filter.Page - 1) * filter.Limit

	err := query.
		Order("created_at DESC").
		Limit(filter.Limit).
		Offset(offset).
		Find(&logs).
		Error

	return logs, total, err
}

func (r *ActivityLogRepository) FindByUserID(
	userID uint,
	page int,
	limit int,
) ([]models.ActivityLog, int64, error) {

	var logs []models.ActivityLog
	var total int64

	query := r.DB.
		Model(&models.ActivityLog{}).
		Where("user_id = ?", userID).
		Preload("User")

	query.Count(&total)

	offset := (page - 1) * limit

	err := query.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&logs).
		Error

	return logs, total, err
}