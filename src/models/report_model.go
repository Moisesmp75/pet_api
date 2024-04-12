package models

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	ID             uint64 `gorm:"primaryKey;autoIncrement"`
	ReporterUserID uint64 `gorm:"not null"`
	ReporterUser   User   `gorm:"foreignKey:ReporterUserID"`
	ReportedUserID uint64 `gorm:"not null"`
	ReportedUser   User   `gorm:"foreignKey:ReportedUserID"`
	Description    string `gorm:"size:400"`
}
