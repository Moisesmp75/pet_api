package models

import "gorm.io/gorm"

type PetBehavior struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	PetID       uint64 `gorm:"not null"`
	Temper      string `gorm:"size:25"`
	Habit       string `gorm:"size:25"`
	Personality string `gorm:"size:25"`
}
