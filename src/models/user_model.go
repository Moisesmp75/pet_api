package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       		uint   `gorm:"primaryKey;autoIncrement"`
	Name 		 		string `gorm:"not null"`
	LastName 		string `gorm:"not null"`
	PhoneNumber string `gorm:"unique; not null"`
	Email    		string `gorm:"unique; not null"`
	Password 		string `gorm:"not null"`
	Pets     		[]Pet  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserID;references:ID"`
	RoleID   		uint   `gorm:"not null"`
	Role     		Role   `gorm:"foreignKey:RoleID"`
}
