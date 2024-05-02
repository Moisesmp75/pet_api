package request

type UpdatePetRequest struct {
	Name        string  `json:"name"`
	Breed       string  `json:"breed"`
	BornDate    string  `json:"born_date" validate:"omitempty,datetime=2006/01/02"`
	Description string  `json:"description"`
	Height      float32 `json:"height"`
	Gender      string  `json:"gender"`
	Color       string  `json:"color"`
	Weight      float32 `json:"weight"`
	Location    string  `json:"location"`
	Adopted     *bool   `json:"adopted"`
	PetTypeId   uint64  `json:"pet_type_id"`
}
