package repositories

import (
	"pet_api/src/database"
	"pet_api/src/models"
)

func CreatePetImage(newPetImage models.PetImage) (models.PetImage, error) {
	if err := database.DB.Model(&models.PetImage{}).Create(&newPetImage).Error; err != nil {
		return models.PetImage{}, err
	}
	return newPetImage, nil
}

func CreatePetImages(newPetImage []models.PetImage) ([]models.PetImage, error) {
	if err := database.DB.Model(&models.PetImage{}).Create(&newPetImage).Error; err != nil {
		return []models.PetImage{}, err
	}
	return newPetImage, nil
}