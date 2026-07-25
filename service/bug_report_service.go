package service

import (
	"errors"

	"github.com/dewialvi/digital-channel-monitoring/models"
	"github.com/dewialvi/digital-channel-monitoring/repository"
)

type BugReportService struct {
	Repo *repository.BugReportRepository
}

func NewBugReportService(repo *repository.BugReportRepository) *BugReportService {
	return &BugReportService{
		Repo: repo,
	}
}

var ErrBugNotFound = errors.New("bug report tidak ditemukan")

type CreateBugReportInput struct {
	ReportedBy       uint
	Title            string
	Description      string
	Severity         models.Severity
	Priority         models.Priority
	StepsToReproduce string
}

func (s *BugReportService) Create(input CreateBugReportInput) (*models.BugReport, error) {
	bug := &models.BugReport{
		ReportedBy:       input.ReportedBy,
		Title:            input.Title,
		Description:      input.Description,
		Severity:         input.Severity,
		Priority:         input.Priority,
		Status:           models.BugStatusNew,
		StepsToReproduce: input.StepsToReproduce,
	}

	err := s.Repo.Create(bug)
	if err != nil {
		return nil, err
	}

	return bug, nil
}

func (s *BugReportService) GetByID(id uint) (*models.BugReport, error) {
	bug, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, ErrBugNotFound
	}

	return bug, nil
}

func (s *BugReportService) GetAll(filter repository.BugReportFilter) ([]models.BugReport, int64, error) {
	if filter.Page < 1 {
		filter.Page = 1
	}

	if filter.Limit < 1 || filter.Limit > 100 {
		filter.Limit = 10
	}

	return s.Repo.FindAll(filter)
}

func (s *BugReportService) UpdateStatus(id uint, newStatus models.BugStatus) (*models.BugReport, error) {
	bug, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, ErrBugNotFound
	}

	if !isValidStatusTransition(bug.Status, newStatus) {
		return nil, errors.New(
			"perubahan status tidak valid: " +
				string(bug.Status) +
				" -> " +
				string(newStatus),
		)
	}

	bug.Status = newStatus

	err = s.Repo.Update(bug)
	if err != nil {
		return nil, err
	}

	return bug, nil
}

func isValidStatusTransition(current, next models.BugStatus) bool {
	allowed := map[models.BugStatus][]models.BugStatus{
		models.BugStatusNew: {
			models.BugStatusAssigned,
		},
		models.BugStatusAssigned: {
			models.BugStatusInProgress,
		},
		models.BugStatusInProgress: {
			models.BugStatusFixed,
		},
		models.BugStatusFixed: {
			models.BugStatusRetesting,
		},
		models.BugStatusRetesting: {
			models.BugStatusVerified,
			models.BugStatusReopened,
		},
		models.BugStatusVerified: {
			models.BugStatusClosed,
		},
		models.BugStatusReopened: {
			models.BugStatusAssigned,
		},
	}

	for _, allowedStatus := range allowed[current] {
		if allowedStatus == next {
			return true
		}
	}

	return false
}

func (s *BugReportService) Delete(id uint) error {
	_, err := s.Repo.FindByID(id)
	if err != nil {
		return ErrBugNotFound
	}

	return s.Repo.Delete(id)
}
