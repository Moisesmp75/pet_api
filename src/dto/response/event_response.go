package response

import (
	"time"
)

type EventResponse struct {
	ID              uint64
	Title           string
	Description     string
	ImageUrl        string
	Date            time.Time
	ONG             UserResponse
	AllowVolunteers bool
	Participants    []UserResponse
}
