package models

import "gorm.io/gorm"

type BankAccount struct {
	gorm.Model
	ID            uint64 `gorm:"primaryKey;autoIncrement"`
	ONGInfoID     uint64 `gorm:"not null"`
	BankName      string `gorm:"not null; size:25"`
	AccountNumber string `gorm:"not null; size:25"`
	CCI           string `gorm:"not null; size:35"`
}
