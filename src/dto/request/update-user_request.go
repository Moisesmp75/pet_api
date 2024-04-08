package request

type UpdateUserRequest struct {
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Dni         string `json:"dni"`
	Address     string `json:"address"`
	City        string `json:"city"`
}
