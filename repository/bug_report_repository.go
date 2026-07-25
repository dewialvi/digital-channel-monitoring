package repository

import (
	"gorm.io/gorm"

	"github.com/dewialvi/digital-channel-monitoring/models"
)

type BugReportRepository struct {
	DB *gorm.DB
}

func NewBugReportRepository(db *gorm.DB) *BugReportRepository {
	return &BugReportRepository{DB: db}
}

type BugReportFilter struct {
	Search   string
	Severity string
	Priority string
	Status   string
	Page     int
	Limit    int
}

func (r *BugReportRepository) Create(bug *models.BugReport) error {
	return r.DB.Create(bug).Error
}

func (r *BugReportRepository) FindByID(id uint) (*models.BugReport, error) {
	var bug models.BugReport

	err := r.DB.
		Preload("Reporter").
		First(&bug, id).
		Error

	if err != nil {
		return nil, err
	}

	return &bug, nil
}

func (r *BugReportRepository) FindAll(filter BugReportFilter) ([]models.BugReport, int64, error) {
	var bugs []models.BugReport
	var total int64

	query := r.DB.
		Model(&models.BugReport{}).
		Preload("Reporter")

	if filter.Search != "" {
		query = query.Where(
			"title ILIKE ? OR description ILIKE ?",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
	}

	if filter.Severity != "" {
		query = query.Where("severity = ?", filter.Severity)
	}

	if filter.Priority != "" {
		query = query.Where("priority = ?", filter.Priority)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (filter.Page - 1) * filter.Limit

	err := query.
		Order("created_at DESC").
		Limit(filter.Limit).
		Offset(offset).
		Find(&bugs).
		Error

	if err != nil {
		return nil, 0, err
	}

	return bugs, total, nil
}

func (r *BugReportRepository) Update(bug *models.BugReport) error {
	return r.DB.Save(bug).Error
}

func (r *BugReportRepository) Delete(id uint) error {
	return r.DB.Delete(&models.BugReport{}, id).Error
}
