package repositories

import (
	"pet_api/src/database"
	"pet_api/src/models"
)

func CountPets() int64 {
	var total_items int64
	database.DB.Model(&models.Pet{}).Count(&total_items)
	return total_items
}

func CreatePet(newPet models.Pet) (*models.Pet, error) {
	if err := database.DB.Create(&newPet).Error; err != nil {
		return nil, err
	}
	return &newPet, nil
}

func GetPetById(id uint) (*models.Pet, error) {
	var pet models.Pet
	data := database.DB.First(&pet, id)

	if data.Error != nil || data.RowsAffected == 0 {
		return nil, data.Error
	}

	return &pet, nil
}

func GetAllPets(offset, limit int) (*[]models.Pet, error) {
	var pets []models.Pet
	data := database.DB.Offset(offset).Limit(limit).Find(&pets)
	if data.Error != nil {
		return nil, data.Error
	}
	return &pets, nil
}
