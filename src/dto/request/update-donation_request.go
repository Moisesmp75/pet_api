package request

type UpdateDonationRequest struct {
	Received bool `json:"received" validate:"required,boolean"`
}

type UpdateDonationMoneyRequest struct {
	Received bool `json:"received" validate:"required,boolean"`
}
