package response

type UserResponse struct {
	ID          uint64        `json:"id"`
	Name        string        `json:"name"`
	LastName    string        `json:"last_name"`
	UserName    string        `json:"user_name"`
	PhoneNumber string        `json:"phone_number"`
	Dni         string        `json:"dni"`
	Address     string        `json:"address"`
	City        string        `json:"city"`
	Email       string        `json:"email"`
	Role        string        `json:"role"`
	Pets        []PetResponse `json:"pets,omitempty"`
	ImageUrl    string        `json:"image_url"`
}
