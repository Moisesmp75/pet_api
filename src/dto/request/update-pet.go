package request

type UpdatePetRequest struct {
	Name        string  `json:"name"`
	Breed       string  `json:"breed"`
	BornDate    string  `json:"born_date"`
	Description string  `json:"description"`
	Height      float32 `json:"height"`
	Gender      string  `json:"gender"`
	Color       string  `json:"color"`
	Weight      float32 `json:"weight"`
	Location    string  `json:"location"`
	PetTypeId   uint64  `json:"pet_type_id"`
}
