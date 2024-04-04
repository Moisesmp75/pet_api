package response

type UserResponse struct {
	ID       uint          `json:"id"`
	Username string        `json:"username"`
	Email    string        `json:"email"`
	Pets     []PetResponse `json:"pets,omitempty"`
}