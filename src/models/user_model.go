package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Username    string `gorm:"size:25"`
	Name        string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	PhoneNumber string `gorm:"unique; not null"`
	Dni         string `gorm:"size:9; unique; not null"`
	Address     string `gorm:"size:50"`
	City        string `gorm:"size:25"`
	Email       string `gorm:"unique; not null"`
	Password    string `gorm:"not null"`
	Pets        []Pet  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserID;references:ID"`
	RoleID      uint   `gorm:"not null"`
	Role        Role   `gorm:"foreignKey:RoleID"`
	ImageUrl    string `gorm:"type:LONGTEXT"`
}
