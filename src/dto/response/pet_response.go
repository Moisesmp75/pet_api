package response

type PetResponse struct {
	ID          uint64              `json:"id"`
	Name        string              `json:"name"`
	Breed       string              `json:"breed"`
	Age         int                 `json:"age"`
	Description string              `json:"description"`
	Height      float32             `json:"height"`
	Gender      string              `json:"gender"`
	Color       string              `json:"color"`
	Weight      float32             `json:"weight"`
	Adopted     bool                `json:"adppted"`
	User        *UserResponse       `json:"user,omitempty"`
	Location    string              `json:"location"`
	PetType     string              `json:"pet_type"`
	Images      string              `json:"images"`
	Behavior    PetBehaviorResponse `json:"behavior"`
}
