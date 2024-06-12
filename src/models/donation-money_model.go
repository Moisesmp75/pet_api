package models

import "gorm.io/gorm"

type DonationMoney struct {
	gorm.Model
	ID       uint64  `gorm:"primaryKey;autoIncrement"`
	UserID   uint64  `gorm:"not null"`
	User     User    `gorm:"foreignKey:UserID"`
	OngID    uint64  `gorm:"not null"`
	Ong      User    `gorm:"foreignKey:OngID"`
	Amount   float32 `gorm:"not null"`
	Received bool    `gorm:"default:false"`
}
