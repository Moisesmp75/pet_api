package response

type UserResponse struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	LastName    string        `json:"last_name"`
	PhoneNumber string        `json:"phone_number"`
	Dni         string        `json:"dni"`
	Address     string        `json:"address"`
	City        string        `json:"city"`
	Email       string        `json:"email"`
	Role        string        `json:"role"`
	Pets        []PetResponse `json:"pets,omitempty"`
	ImageUrl    string        `json:"image_url"`
}
