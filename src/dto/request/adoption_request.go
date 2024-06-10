package request

type AdoptionRequest struct {
	PetID        uint64 `json:"pet_id" validate:"required"`
	AdoptionDate string `json:"adoption_date" validate:"required,datetime=2006/01/02 15:04:05"`
	Comment      string `json:"comment" validate:"required,gt=0"`
}
