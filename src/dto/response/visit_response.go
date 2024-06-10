package response

import "time"

type VisitResponse struct {
	ID    uint64       `json:"id"`
	Pet   PetResponse  `json:"pet"`
	User  UserResponse `json:"user"`
	Date  time.Time    `json:"date"`
	State string       `json:"state"`
}
