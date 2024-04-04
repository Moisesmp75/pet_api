package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique; not null"`
	Password string `gorm:"not null"`
	Pets     []Pet 	`gorm:"foreignKey:UserID;references:ID"`
}
