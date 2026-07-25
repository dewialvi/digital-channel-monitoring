package service

import (
	"github.com/dewialvi/digital-channel-monitoring/models"
	"github.com/dewialvi/digital-channel-monitoring/repository"
)

type ActivityLogService struct {
	Repo *repository.ActivityLogRepository
}

func NewActivityLogService(
	repo *repository.ActivityLogRepository,
) *ActivityLogService {
	return &ActivityLogService{
		Repo: repo,
	}
}

func (s *ActivityLogService) Create(
	userID uint,
	action string,
	description string,
) error {

	log := &models.ActivityLog{
		UserID:      userID,
		Action:      action,
		Description: description,
	}

	return s.Repo.Create(log)
}

func (s *ActivityLogService) GetAll(
	page int,
	limit int,
) ([]models.ActivityLog, int64, error) {

	filter := repository.ActivityLogFilter{
		Page:  page,
		Limit: limit,
	}

	return s.Repo.FindAll(filter)
}
func (s *ActivityLogService) Log(
	userID uint,
	action string,
	description string,
) error {

	log := &models.ActivityLog{
		UserID:      userID,
		Action:      action,
		Description: description,
	}

	return s.Repo.Create(log)
}