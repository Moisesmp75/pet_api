package response

import "time"

type DonationProductResponse struct {
	ID           uint64            `json:"id"`
	User         UserResponse      `json:"user"`
	Ong          UserResponse      `json:"ong"`
	Products     []ProductResponse `json:"products"`
	DonationDate time.Time         `json:"donation_date"`
}
