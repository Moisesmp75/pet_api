package request

type UserRequest struct {
	Name        string `json:"name" validate:"required,gt=0"`
	UserName    string `json:"user_name"`
	LastName    string `json:"last_name" validate:"required,gt=0"`
	PhoneNumber string `json:"phone_number" validate:"required,gt=0"`
	Dni         string `json:"dni" validate:"required,gt=0"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,gt=0"`
	RoleID      uint   `json:"role_id" validate:"required"`
}
