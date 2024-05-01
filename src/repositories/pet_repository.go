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
	data := database.DB.Model(&models.Pet{}).Preload("PetType").Preload("User").Preload("Images")
	data = data.Preload("User.Role").First(&pet, id)

	if data.Error != nil || data.RowsAffected == 0 {
		return models.Pet{}, fmt.Errorf("pet with id '%d' not found", id)
	}

	// if err := database.DB.Preload("Role").First(&pet.User, pet.UserID).Error; err != nil {
	// 	return models.Pet{}, err
	// }
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

	data := query.Offset(offset).Limit(limit).Preload("User").Preload("Images")

	data = data.Preload("User.Role").Preload("PetType").Find(&pets)
	if data.Error != nil {
		return nil, data.Error
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

func DeletePet(id uint64) (models.Pet, error) {
	pet, err := GetPetById(id)
	if err != nil {
		return models.Pet{}, err
	}
	operation := database.DB.Select("Images").Delete(&pet)

	if operation.Error != nil || operation.RowsAffected == 0 {
		return models.Pet{}, err
	}
	return pet, nil
}

func DeletePets(pets []models.Pet) ([]models.Pet, error) {
	if len(pets) == 0 {
		return []models.Pet{}, nil
	}
	operation := database.DB.Select("Images").Delete(&pets)
	if operation.Error != nil || operation.RowsAffected == 0 {
		return []models.Pet{}, operation.Error
	}
	return pets, nil
}
