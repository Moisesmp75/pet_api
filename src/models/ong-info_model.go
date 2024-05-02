package models

import "gorm.io/gorm"

type ONGInfo struct {
	gorm.Model
	ID           uint64        `gorm:"primaryKey;autoIncrement"`
	UserID       uint64        `gorm:"unique; not null"`
	Description  string        `gorm:"size:400"`
	BankAccounts []BankAccount `gorm:"constraint:OnDelete:CASCADE;foreignKey:ONGInfoID"`
}
