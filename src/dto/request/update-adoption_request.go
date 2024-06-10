package request

type UpdateAdoptionRequest struct {
	State string `json:"state" validate:"required,gt=0"`
}
