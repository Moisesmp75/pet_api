package response

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Role  string `json:"role"`
}
