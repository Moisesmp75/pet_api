package models

import (
	"time"

	"gorm.io/gorm"
)

type Adoption struct {
	gorm.Model
	ID           uint64    `gorm:"primaryKey;autoIncrement"`
	PetID        uint64    `gorm:"not null"`
	Pet          Pet       `gorm:"foreignKey:PetID"`
	UserID       uint64    `gorm:"not null"`
	User         User      `gorm:"foreignKey:UserID"`
	AdoptionDate time.Time `gorm:"not null"`
	Comment      string    `gorm:"size:150"`
}
