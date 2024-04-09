package models

import "gorm.io/gorm"

type PetType struct {
	gorm.Model
	ID   uint64 `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"unique;not null"`
}
