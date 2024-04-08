package models

import "gorm.io/gorm"

type PetImage struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	URL     string `gorm:"not null"`
	PetID   uint   `gorm:"not null"`
}