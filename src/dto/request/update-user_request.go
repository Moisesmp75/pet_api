package request

type UpdateUserRequest struct {
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
