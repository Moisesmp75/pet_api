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
		AdoptionDate: helpers.ParseDateTime(req.AdoptionDate),
		Comment:      req.Comment,
		State:        "Pending",
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
		State:           adoption.State,
	}
}

func AdoptionModelsToResponse(adoptions []models.Adoption) []response.AdoptionResponse {
	resp := make([]response.AdoptionResponse, len(adoptions))

	for i, v := range adoptions {
		resp[i] = AdoptionModelToResponse(v)
	}

	return resp
}

func UpdateAdoptionRequestToModel(req request.UpdateAdoptionRequest, adoption models.Adoption) models.Adoption {
	if req.State != "" {
		adoption.State = req.State
	}
	return adoption
}
