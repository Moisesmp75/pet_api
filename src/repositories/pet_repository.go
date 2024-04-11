package repositories

import (
	"fmt"
	"pet_api/src/database"
	"pet_api/src/models"
)

func CountPets(breed, color string) int64 {
	var total_items int64
	query := database.DB.Model(&models.Pet{})

	if breed != "" {
		query = query.Where("breed = ?", breed)
	}
	if color != "" {
		query = query.Where("color = ?", color)
	}

	if err := query.Count(&total_items).Error; err != nil {
		return 0
	}
	return total_items
}

func CreatePet(newPet models.Pet) (models.Pet, error) {
	if err := database.DB.Model(&models.Pet{}).Create(&newPet).Error; err != nil {
		return models.Pet{}, err
	}
	return newPet, nil
}

func GetPetById(id uint64) (models.Pet, error) {
	var pet models.Pet
	data := database.DB.Model(&models.Pet{}).Preload("PetType").Preload("User").Preload("Images").First(&pet, id)

	if data.Error != nil || data.RowsAffected == 0 {
		return models.Pet{}, fmt.Errorf("pet with id '%d' not found", id)
	}

	if err := database.DB.Preload("Role").First(&pet.User, pet.UserID).Error; err != nil {
		return models.Pet{}, err
	}
	return pet, nil
}

func GetAllPets(offset, limit int, breed, color string) ([]models.Pet, error) {
	var pets []models.Pet
	query := database.DB.Model(&models.Pet{})

	if breed != "" {
		query = query.Where("breed = ?", breed)
	}

	if color != "" {
		query = query.Where("color = ?", color)
	}

	data := query.Offset(offset).Limit(limit).Preload("User").Preload("Images").Find(&pets)
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

func UpdatePet(pet models.Pet) (models.Pet, error) {
	data := database.DB.Model(&models.Pet{}).Where("id = ?", pet.ID).Updates(pet)
	if data.RowsAffected == 0 || data.Error != nil {
		return models.Pet{}, data.Error
	}
	return pet, nil
}
