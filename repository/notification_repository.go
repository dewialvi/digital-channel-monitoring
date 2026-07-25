package repository

import (
	"gorm.io/gorm"

	"github.com/dewialvi/digital-channel-monitoring/models"
)

type NotificationRepository struct {
	DB *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{DB: db}
}

func (r *NotificationRepository) Create(n *models.Notification) error {
	return r.DB.Create(n).Error
}

func (r *NotificationRepository) FindByUserID(
	userID uint,
	page int,
	limit int,
) ([]models.Notification, int64, error) {

	var notifs []models.Notification
	var total int64

	query := r.DB.
		Model(&models.Notification{}).
		Where("user_id = ?", userID)

	query.Count(&total)

	offset := (page - 1) * limit

	err := query.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&notifs).
		Error

	return notifs, total, err
}

func (r *NotificationRepository) MarkAsRead(
	id uint,
	userID uint,
) error {
	return r.DB.
		Model(&models.Notification{}).
		Where("id = ? AND user_id = ?", id, userID).
		Update("is_read", true).
		Error
}

func (r *NotificationRepository) CountUnread(
	userID uint,
) (int64, error) {

	var count int64

	err := r.DB.
		Model(&models.Notification{}).
		Where(
			"user_id = ? AND is_read = ?",
			userID,
			false,
		).
		Count(&count).
		Error

	return count, err
}