package mapper

import (
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/models"
)

func PetBehaviorModelToResponse(petBehavior models.PetBehavior) response.PetBehaviorResponse {
	return response.PetBehaviorResponse{
		Temper:      petBehavior.Temper,
		Habit:       petBehavior.Habit,
		Personality: petBehavior.Personality,
	}
}

func PetBehaviorRequestToModel(petBehaviorRequest request.PetBehaviorRequest) models.PetBehavior {
	return models.PetBehavior{
		Temper:      petBehaviorRequest.Temper,
		Habit:       petBehaviorRequest.Habit,
		Personality: petBehaviorRequest.Personality,
	}
}
