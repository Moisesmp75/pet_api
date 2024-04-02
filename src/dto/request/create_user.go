package request

type UserRequest struct {
	Username string `json:"username" validate:"required,gt=0"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gt=0"`
}