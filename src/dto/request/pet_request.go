package request

import "time"

type PetRequest struct {
	Name        string    `json:"name" validate:"required,gt=0"`
	Breed       string    `json:"breed" validate:"required,gt=0"`
	BornDate    time.Time `json:"born_date" validate:"required,datetime"`
	Description string    `json:"description"`
	Size        float32   `json:"size" validate:"required,gt=0"`
	Gender      string    `json:"gender" validate:"len=1"`
	Color       string    `json:"color" validate:"required"`
	Weight      float32   `json:"weight" validate:"required,gt=0"`
	UserID      uint      `json:"user_id" validate:"required,gt=0"`
}