package response

import (
	"time"
)

type PetResponse struct {
	ID          uint          `json:"id"`
	Breed       string        `json:"breed"`
	BornDate    time.Time     `json:"born_date"`
	Description string        `json:"description"`
	Height      float32       `json:"height"`
	Gender      string        `json:"gender"`
	Color       string        `json:"color"`
	Weight      float32       `json:"weight"`
	User        *UserResponse `json:"user,omitempty"`
	Location    string        `json:"location"`
}
