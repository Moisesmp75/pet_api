package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"size:20;not null" validate:"required"`
	Price       float32 `gorm:"not null" validate:"required"`
	Description string  `gorm:"size:100"`
}