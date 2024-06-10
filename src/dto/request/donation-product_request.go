package request

type DonationProductRequest struct {
	OngId    uint64           `json:"ong_id" validate:"required"`
	Products []ProductRequest `json:"products" validate:"required"`
}
