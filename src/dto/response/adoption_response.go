package response

import "time"

type AdoptionResponse struct {
	ID              uint64       `json:"id"`
	User            UserResponse `json:"user"`
	Pet             PetResponse  `json:"pet"`
	ApplicationDate time.Time    `json:"application_date"`
	AdoptionDate    time.Time    `json:"adoption_date"`
	Comment         string       `json:"comment"`
}
