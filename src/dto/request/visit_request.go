package request

type VisitRequest struct {
	PetID uint64 `json:"pet_id" validate:"required"`
	Date  string `json:"date" validate:"required,datetime=2006/01/02 15:04:05"`
}
