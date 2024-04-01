package request

type ProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}
