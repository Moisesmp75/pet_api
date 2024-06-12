package request

type DonationMoneyRequest struct {
	OngID  uint64  `json:"ong_id" validate:"required"`
	Amount float32 `json:"amount" validate:"required"`
}
