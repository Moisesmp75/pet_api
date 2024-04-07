package models

import "gorm.io/gorm"

type PetImage struct {
	gorm.Model
	URL     string `gorm:"not null"`
	PetID   uint   `gorm:"not null"`
}