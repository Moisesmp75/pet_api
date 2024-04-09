package request

type LoginRequest struct {
	Identity string `json:"identity" validate:"required"`
	Password string `json:"password" validate:"required,gt=0"`
	RoleID   uint64 `json:"role_id" validate:"required"`
}