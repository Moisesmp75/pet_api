package response

import "time"

type DonationMoneyResponse struct {
	ID           uint64       `json:"id"`
	User         UserResponse `json:"user"`
	Ong          UserResponse `json:"ong"`
	Amount       float32      `json:"amount"`
	DonationDate time.Time    `json:"donation_date"`
	Received     bool         `json:"received"`
}
