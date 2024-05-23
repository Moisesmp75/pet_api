package models

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"size:20;not null"`
	Breed       string    `gorm:"size:30;not null"`
	BornDate    time.Time `gorm:"not null"`
	Description string    `gorm:"size:500"`
	Height      float32   `gorm:"not null"`
	Gender      string    `gorm:"size:1"`
	Color       string    `gorm:"size:40"`
	Weight      float32   `gorm:"not null"`
	Adopted     bool      `gorm:"not null;default:false"`
	UserID      uint64    `gorm:"not null"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`
	Image       PetImage  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PetID"`
	Location    string    `gorm:"size:50"`
	PetTypeId   uint64    `gorm:"not null"`
	PetType     PetType   `gorm:"foreignKey:PetTypeId"`
}
