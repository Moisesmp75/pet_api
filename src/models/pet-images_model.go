package models

import "gorm.io/gorm"

type PetImage struct {
	gorm.Model
	ID      uint64 `gorm:"primaryKey;autoIncrement"`
	URL     string `gorm:"not null"`
	PetID   uint64 `gorm:"not null"`
}