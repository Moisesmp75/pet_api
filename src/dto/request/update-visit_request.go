package request

type UpdateVisitRequest struct {
	State string `json:"state" validate:"required,gt=0"`
}
