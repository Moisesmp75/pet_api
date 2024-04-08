package mapper

import "pet_api/src/models"

func PetImagesModelToResponse(petImages []models.PetImage) []string {
	urls := make([]string, len(petImages))

	for i, v := range petImages {
		urls[i] = v.URL
	}

	return urls
}