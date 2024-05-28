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

func UpdatePetImage(petImage models.PetImage) (models.PetImage, error) {
	data := database.DB.Model(&models.PetImage{}).Where("id = ?", petImage.ID).Updates(petImage)
	if data.RowsAffected == 0 || data.Error != nil {
		return models.PetImage{}, data.Error
	}
	return petImage, nil
}
