package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID              uint64 `gorm:"primaryKey;autoIncrement"`
	Title           string `gorm:"size:80;not null"`
	Description     string `gorm:"size:1000;not null"`
	ImageUrl        string `gorm:"type:LONGTEXT"`
	ONGID           uint64 `gorm:"not null"`
	ONG             User   `gorm:"foreignKey:ONGID"`
	AllowVolunteers bool   `gorm:"not null"`
	Participants    []User `gorm:"many2many:event_participants;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
