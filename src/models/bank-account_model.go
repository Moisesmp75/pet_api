package models

import "gorm.io/gorm"

type BankAccount struct {
	gorm.Model
	ONGInfoID    uint64 `gorm:"not null"`
	BankName     string `gorm:"not null; size:25"`
	AccountNumer string `gorm:"not null; size:25"`
	CCI          string `gorm:"not null; size:35"`
}
