package repository

import (
	"time"

	"gorm.io/gorm"

	"github.com/dewialvi/digital-channel-monitoring/models"
)

type APIMonitorRepository struct {
	DB *gorm.DB
}

func NewAPIMonitorRepository(db *gorm.DB) *APIMonitorRepository {
	return &APIMonitorRepository{
		DB: db,
	}
}

func (r *APIMonitorRepository) Create(m *models.APIMonitor) error {
	return r.DB.Create(m).Error
}

type APIMonitorFilter struct {
	Endpoint   string
	StatusCode int
	Page       int
	Limit      int
}

func (r *APIMonitorRepository) FindAll(
	filter APIMonitorFilter,
) ([]models.APIMonitor, int64, error) {

	var logs []models.APIMonitor
	var total int64

	query := r.DB.Model(&models.APIMonitor{})

	if filter.Endpoint != "" {
		query = query.Where(
			"endpoint ILIKE ?",
			"%"+filter.Endpoint+"%",
		)
	}

	if filter.StatusCode != 0 {
		query = query.Where(
			"status_code = ?",
			filter.StatusCode,
		)
	}

	query.Count(&total)

	offset := (filter.Page - 1) * filter.Limit

	err := query.
		Order("checked_at DESC").
		Limit(filter.Limit).
		Offset(offset).
		Find(&logs).
		Error

	return logs, total, err
}

type StatsSummary struct {
	TotalRequests    int64   `json:"total_requests"`
	ErrorCount       int64   `json:"error_count"`
	ErrorRatePercent float64 `json:"error_rate_percent"`
	AvgResponseTime  float64 `json:"avg_response_time_ms"`
}

func (r *APIMonitorRepository) GetStatsSummary(
	since time.Time,
) (*StatsSummary, error) {

	var summary StatsSummary

	err := r.DB.
		Model(&models.APIMonitor{}).
		Where("checked_at >= ?", since).
		Count(&summary.TotalRequests).
		Error

	if err != nil {
		return nil, err
	}

	err = r.DB.
		Model(&models.APIMonitor{}).
		Where(
			"checked_at >= ? AND status_code >= 500",
			since,
		).
		Count(&summary.ErrorCount).
		Error

	if err != nil {
		return nil, err
	}

	if summary.TotalRequests > 0 {
		summary.ErrorRatePercent =
			float64(summary.ErrorCount) /
				float64(summary.TotalRequests) *
				100
	}

	var avgResult struct {
		Avg float64
	}

	err = r.DB.
		Model(&models.APIMonitor{}).
		Select("AVG(response_time_ms) as avg").
		Where("checked_at >= ?", since).
		Scan(&avgResult).
		Error

	if err != nil {
		return nil, err
	}

	summary.AvgResponseTime = avgResult.Avg

	return &summary, nil
}