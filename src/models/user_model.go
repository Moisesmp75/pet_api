package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null" validates:"required"`
	Email    string `gorm:"unique; not null" validates:"required"`
	Password string `gorm:"not null" validates:"required"`
}
