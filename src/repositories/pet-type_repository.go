package repositories

import (
	"fmt"
	"pet_api/src/database"
	"pet_api/src/models"
)

func CountPetTypes() int64 {
	var total_items int64
	if err := database.DB.Model(&models.PetType{}).Count(&total_items).Error; err != nil {
		return 0
	}
	return total_items
}

func GetPetTypeById(id uint64) (models.PetType, error) {
	var petType models.PetType
	data := database.DB.Model(&models.PetType{}).First(&petType, id)

	if data.Error != nil || data.RowsAffected == 0 {
		return models.PetType{}, fmt.Errorf("pet_type with id '%d' not found", id)
	}

	return petType, nil
}

func GetAllPetTypes() ([]models.PetType, error) {
	var petTypes []models.PetType

	data := database.DB.Model(&models.PetType{}).Find(&petTypes)
	if data.Error != nil {
		return nil, data.Error
	}

	return petTypes, nil
}
