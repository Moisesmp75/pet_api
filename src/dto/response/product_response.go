package response

type ProductResponse struct {
	Unit     float32 `json:"unit"`
	Quantity int     `json:"quantity"`
	Name     string  `json:"name"`
}
