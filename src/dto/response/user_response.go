package response

type UserResponse struct {
	ID       uint          `json:"id"`
	Username string        `json:"username"`
	Email    string        `json:"email"`
	Role     string        `json:"role"`
	Pets     []PetResponse `json:"pets,omitempty"`
}
