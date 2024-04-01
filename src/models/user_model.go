package models

type User struct {
	// gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null" validate:"required"`
	Email    string `gorm:"unique; not null" validate:"required"`
	Password string `gorm:"not null" validate:"required"`
}
