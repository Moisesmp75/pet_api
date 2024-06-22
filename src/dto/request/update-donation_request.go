package request

type UpdateDonationRequest struct {
	Received bool `json:"received" validate:"boolean"`
}

type UpdateDonationMoneyRequest struct {
	Received bool `json:"received" validate:"boolean"`
}
