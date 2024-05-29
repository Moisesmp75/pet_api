package request

type PetBehaviorRequest struct {
	Temper      string `json:"temper" validate:"required,gt=0"`
	Habit       string `json:"habit" validate:"required,gt=0"`
	Personality string `json:"personality" validate:"required,gt=0"`
}
