package response

import (
	"time"
)

type EventResponse struct {
	ID              uint64         `json:"id"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	ImageUrl        string         `json:"image_url"`
	PublicationDate time.Time      `json:"publication_date"`
	ONG             UserResponse   `json:"ong"`
	AllowVolunteers bool           `json:"allow_volunteers"`
	Participants    []UserResponse `json:"participants"`
}
