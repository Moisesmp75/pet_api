package mapper

import (
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/helpers"
	"pet_api/src/models"
)

func AdoptionRequestToModel(req request.AdoptionRequest) models.Adoption {
	return models.Adoption{
		PetID:        req.PetID,
		UserID:       req.UserID,
		AdoptionDate: helpers.ParseDate(req.AdoptionDate),
		Comment:      req.Comment,
	}
}

func AdoptionModelToResponse(adoption models.Adoption) response.AdoptionResponse {
	return response.AdoptionResponse{
		ID:              adoption.ID,
		User:            *OnlyUserModelToResponse(adoption.User),
		Pet:             PetModelToResponse(adoption.Pet),
		ApplicationDate: adoption.CreatedAt,
		AdoptionDate:    adoption.AdoptionDate,
		Comment:         adoption.Comment,
	}
}
