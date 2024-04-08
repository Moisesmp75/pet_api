package response

type PetResponse struct {
	ID          uint          `json:"id"`
	Breed       string        `json:"breed"`
	Age         int           `json:"age"`
	Description string        `json:"description"`
	Height      float32       `json:"height"`
	Gender      string        `json:"gender"`
	Color       string        `json:"color"`
	Weight      float32       `json:"weight"`
	User        *UserResponse `json:"user,omitempty"`
	Location    string        `json:"location"`
	PetType     string        `json:"pet_type"`
}
