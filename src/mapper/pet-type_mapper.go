package mapper

import (
	"pet_api/src/dto/response"
	"pet_api/src/models"
)

func PetTypeModelToResponse(petType models.PetType) response.PetTypeResponse {
	return response.PetTypeResponse{
		ID:   petType.ID,
		Name: petType.Name,
	}
}

func PetTypeModelsToResponse(petTypes []models.PetType) []response.PetTypeResponse {
	resp := make([]response.PetTypeResponse, len(petTypes))

	for i, v := range petTypes {
		resp[i] = PetTypeModelToResponse(v)
	}

	return resp
}
