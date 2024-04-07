package request

type PasswordResetRequest struct {
	Email string `json:"email" validate:"required,email"`
}
