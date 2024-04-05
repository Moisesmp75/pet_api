package response

import "time"

type LoginResponse struct {
	Email string   `json:"email"`
	Token string   `json:"token"`
	Role  string   `json:"role"`
	Iat   time.Time `json:"iat"`
	Exp   time.Time    `json:"exp"`
}
