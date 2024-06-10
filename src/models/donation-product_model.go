package models

import "gorm.io/gorm"

type DonationProduct struct {
	gorm.Model
	ID       uint64    `gorm:"primaryKey;autoIncrement"`
	UserID   uint64    `gorm:"not null"`
	User     User      `gorm:"foreignKey:UserID"`
	OngID    uint64    `gorm:"not null"`
	Ong      User      `gorm:"foreignKey:OngID"`
	Products []Product `gorm:"constraint:OnDelete:CASCADE;foreignKey:DonationProductID;references:ID"`
}
