package request

import "mime/multipart"

// type VisitRequest struct {
// 	PetID uint64 `json:"pet_id" validate:"required"`
// 	Date  string `json:"date" validate:"required,datetime=2006/01/02 15:04:05"`
// }

type EventRequest struct {
	Title           string                `json:"title" validate:"required,min=2,max=80"`
	Description     string                `json:"description" validate:"required,max=1000"`
	AllowVolunteers bool                  `json:"allow_volunteers"`
	Image           *multipart.FileHeader `validate:"required"`
}
