package repositories

import (
	"fmt"
	"pet_api/src/database"
	"pet_api/src/models"
)

func CountPets() int64 {
	var total_items int64
	database.DB.Model(&models.Pet{}).Count(&total_items)
	return total_items
}

func CreatePet(newPet models.Pet) (models.Pet, error) {
	if err := database.DB.Model(&models.Pet{}).Create(&newPet).Error; err != nil {
		return models.Pet{}, err
	}
	return newPet, nil
}

func GetPetById(id uint) (models.Pet, error) {
	var pet models.Pet
	data := database.DB.Model(&models.Pet{}).Preload("User").First(&pet, id)

	if data.Error != nil || data.RowsAffected == 0 {
		return models.Pet{}, fmt.Errorf("pet with id '%d' not found", id)
	}

	if err := database.DB.Preload("Role").First(&pet.User, pet.UserID).Error; err != nil {
		return models.Pet{}, err
	}
	return pet, nil
}

func GetAllPets(offset, limit int) ([]models.Pet, error) {
	var pets []models.Pet
	data := database.DB.Model(&models.Pet{}).Offset(offset).Limit(limit).Preload("User").Find(&pets)
	if data.Error != nil {
		return nil, data.Error
	}

	for i := range pets {
		if err := database.DB.Preload("Role").First(&pets[i].User, pets[i].UserID).Error; err != nil {
			return nil, err
		}
	}
	return pets, nil
}
