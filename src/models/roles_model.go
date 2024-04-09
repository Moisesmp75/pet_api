package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"unique;not null"`
	Description string `gorm:"size:100"`
}