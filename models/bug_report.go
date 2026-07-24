package models

import "time"

type Severity string
type Priority string
type BugStatus string

const (
	SeverityCritical Severity = "critical"
	SeverityHigh     Severity = "high"
	SeverityMedium   Severity = "medium"
	SeverityLow      Severity = "low"

	PriorityP1 Priority = "P1"
	PriorityP2 Priority = "P2"
	PriorityP3 Priority = "P3"
	PriorityP4 Priority = "P4"

	BugStatusNew        BugStatus = "new"
	BugStatusAssigned   BugStatus = "assigned"
	BugStatusInProgress BugStatus = "in_progress"
	BugStatusFixed      BugStatus = "fixed"
	BugStatusRetesting  BugStatus = "retesting"
	BugStatusVerified   BugStatus = "verified"
	BugStatusClosed     BugStatus = "closed"
	BugStatusReopened   BugStatus = "reopened"
)

type BugReport struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	ReportedBy       uint      `json:"reported_by" gorm:"not null"`
	Reporter         User      `json:"reporter" gorm:"foreignKey:ReportedBy"`
	Title            string    `json:"title" gorm:"not null"`
	Description      string    `json:"description" gorm:"type:text"`
	Severity         Severity  `json:"severity" gorm:"type:varchar(20);not null"`
	Priority         Priority  `json:"priority" gorm:"type:varchar(10);not null"`
	Status           BugStatus `json:"status" gorm:"type:varchar(20);not null;default:'new'"`
	StepsToReproduce string    `json:"steps_to_reproduce" gorm:"type:text"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (BugReport) TableName() string {
	return "bug_reports"
}