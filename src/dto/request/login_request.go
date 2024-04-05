package request

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gt=0"`
	RoleID   uint   `json:"role_id" validate:"required"`
}