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
		Height:      req.Height,
		Gender:      req.Gender,
		Color:       req.Color,
		Weight:      req.Weight,
		UserID:      0,
		Location:    req.Location,
		Behavior:    PetBehaviorRequestToModel(req.Behavior),
	}
}

func PetModelToResponse(pet models.Pet) response.PetResponse {
	return response.PetResponse{
		ID:          pet.ID,
		Name:        pet.Name,
		Breed:       pet.Breed,
		Age:         helpers.CalculateAge(pet.BornDate),
		Description: pet.Description,
		Height:      pet.Height,
		Gender:      pet.Gender,
		Color:       pet.Color,
		Weight:      pet.Weight,
		Adopted:     pet.Adopted,
		User:        OnlyUserModelToResponse(pet.User),
		Location:    pet.Location,
		PetType:     pet.PetType.Name,
		Images:      PetImageModelToResponse(pet.Image),
		Behavior:    PetBehaviorModelToResponse(pet.Behavior),
	}
}

func OnlyPetModelToResponse(pet models.Pet) response.PetResponse {
	return response.PetResponse{
		ID:          pet.ID,
		Name:        pet.Name,
		Breed:       pet.Breed,
		Age:         helpers.CalculateAge(pet.BornDate),
		Description: pet.Description,
		Height:      pet.Height,
		Gender:      pet.Gender,
		Color:       pet.Color,
		Weight:      pet.Weight,
		Adopted:     pet.Adopted,
		User:        nil,
		Location:    pet.Location,
		PetType:     pet.PetType.Name,
		Images:      PetImageModelToResponse(pet.Image),
		Behavior:    PetBehaviorModelToResponse(pet.Behavior),
	}
}

func UpdatePetRequestToModel(req request.UpdatePetRequest, pet models.Pet) models.Pet {
	if req.Name != "" {
		pet.Name = req.Name
	}
	if req.Breed != "" {
		pet.Breed = req.Breed
	}
	if req.BornDate != "" {
		pet.BornDate = helpers.ParseDate(req.BornDate)
	}
	if req.Description != "" {
		pet.Description = req.Description
	}
	if req.Height != 0 {
		pet.Height = req.Height
	}
	if req.Gender != "" {
		pet.Gender = req.Gender
	}
	if req.Color != "" {
		pet.Color = req.Color
	}
	if req.Weight != 0 {
		pet.Weight = req.Weight
	}
	if req.Location != "" {
		pet.Location = req.Location
	}
	if req.Adopted != nil {
		pet.Adopted = *req.Adopted
	}
	return pet
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
