package models

import (
	"time"

	"gorm.io/gorm"
)

type Visit struct {
	gorm.Model
	ID     uint64	   `gorm:"primaryKey;autoIncrement"`
	PetID  uint64 	 `gorm:"not null"`
	Pet 	 Pet  		 `gorm:"foreignKey:PetID"`
	UserID uint64		 `gorm:"not null"`
	User 	 User 		 `gorm:"foreignKey:UserID"`
	Date   time.Time `gorm:"not null"`
}