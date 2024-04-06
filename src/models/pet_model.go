package models

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	Name        string     `gorm:"size:20;not null"`
	Breed       string     `gorm:"size:20;not null"`
	BornDate    time.Time  `gorm:"not null"`
	Description string     `gorm:"size:100"`
	Size        float32    `gorm:"not null"`
	Gender      string     `gorm:"size:1"`
	Color       string     `gorm:"size:40"`
	Weight      float32    `gorm:"not null"`
	UserID      uint64
	User        User 			 `gorm:"foreignKey:UserID"`
	Images      []PetImage `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PetID"`
}
