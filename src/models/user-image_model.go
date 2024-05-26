package models

import "gorm.io/gorm"

type UserImage struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	URL      string `gorm:"not null;type:LONGTEXT"`
	Filename string `gorm:"not null"`
	UserID   uint64 `gorm:"not null"`
}
