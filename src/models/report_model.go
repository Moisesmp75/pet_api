package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	UserID      uint64 `gorm:"not null"`
	User        User   `gorm:"foreignKey:UserID"`
	Description string `gorm:"size:400"`
}
