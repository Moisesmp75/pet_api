package mapper

import (
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/helpers"
	"pet_api/src/models"
)

func PetRequestToModel(req request.PetRequest) models.Pet {
	return models.Pet{
		Name:        req.Name,
		Breed:       req.Breed,
		BornDate:    helpers.ParseDate(req.BornDate),
		Description: req.Description,
		Size:        req.Size,
		Gender:      req.Gender,
		Color:       req.Color,
		Weight:      req.Weight,
		UserID:      req.UserID,
		Location:    req.Location,
	}
}

func PetModelToResponse(pet models.Pet) response.PetResponse {
	return response.PetResponse{
		ID:          pet.ID,
		Breed:       pet.Breed,
		BornDate:    pet.BornDate,
		Description: pet.Description,
		Size:        pet.Size,
		Gender:      pet.Gender,
		Color:       pet.Color,
		Weight:      pet.Weight,
		User:        OnlyUserModelToResponse(pet.User),
		Location:    pet.Location,
	}
}

func OnlyPetModelToResponse(pet models.Pet) response.PetResponse {
	return response.PetResponse{
		ID:          pet.ID,
		Breed:       pet.Breed,
		BornDate:    pet.BornDate,
		Description: pet.Description,
		Size:        pet.Size,
		Gender:      pet.Gender,
		Color:       pet.Color,
		Weight:      pet.Weight,
		User:        nil,
		Location:    pet.Location,
	}
}

func PetsModelsToResponse(pets []models.Pet) []response.PetResponse {
	resp := make([]response.PetResponse, len(pets))

	for i, v := range pets {
		resp[i] = PetModelToResponse(v)
	}

	return resp
}

func OnlyPetsModelsToResponse(pets []models.Pet) []response.PetResponse {
	resp := make([]response.PetResponse, len(pets))

	for i, v := range pets {
		resp[i] = OnlyPetModelToResponse(v)
	}

	return resp
}
