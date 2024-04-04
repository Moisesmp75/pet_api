package request

type PetRequest struct {
	Name        string  `json:"name" validate:"required,gt=0"`
	Breed       string  `json:"breed" validate:"required,gt=0"`
	BornDate    string  `json:"born_date" validate:"required,datetime=2006/01/02"`
	Description string  `json:"description"`
	Size        float32 `json:"size" validate:"required"`
	Gender      string  `json:"gender" validate:"len=1"`
	Color       string  `json:"color" validate:"required"`
	Weight      float32 `json:"weight" validate:"required"`
	UserID      uint    `json:"user_id" validate:"required"`
}
