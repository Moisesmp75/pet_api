package models

import "time"

type Pet struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"size:20;not null"`
	Breed       string    `gorm:"size:20;not null"`
	BornDate    time.Time `gorm:"not null"`
	Description string    `gorm:"size:100"`
	Size        float32   `gorm:"not null"`
	Gender      string    `gorm:"size:1"`
	Color       string    `gorm:"size:40"`
	Weight      float32   `gorm:"not null"`
}
