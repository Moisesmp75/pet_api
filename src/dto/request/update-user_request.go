package request

type UpdateUserRequest struct {
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	PhoneNumber    string `json:"phone_number"`
	UserName       string `json:"user_name"`
	Password       string `json:"password"`
	Email          string `json:"email" validate:"omitempty,email"`
	Address        string `json:"address"`
	City           string `json:"city"`
	MotherLastName string `json:"mother_last_name"`
	About          string `json:"about"`
}
