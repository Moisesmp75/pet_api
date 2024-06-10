package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID                uint64  `gorm:"primaryKey;autoIncrement"`
	Unit              float32 `gorm:"not null"`
	Quantity          int     `gorm:"not null"`
	Name              string  `gorm:"not null;size:20"`
	DonationProductID uint64  `gorm:"not null"`
}
