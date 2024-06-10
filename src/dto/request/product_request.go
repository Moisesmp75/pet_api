package request

type ProductRequest struct {
	Unit     float32 `json:"unit" validate:"required"`
	Quantity int     `json:"quantity" validate:"required"`
	Name     string  `json:"name" validate:"required"`
}
