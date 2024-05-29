package repositories

import (
	"fmt"
	"pet_api/src/database"
	"pet_api/src/models"
)

func CountPets(breed, color, gender, petType string) int64 {
	var total_items int64
	query := database.DB.Model(&models.Pet{}).Joins("INNER JOIN pet_types ON pets.pet_type_id = pet_types.id")

	if breed != "" {
		query = query.Where("breed = ?", breed)
	}
	if color != "" {
		query = query.Where("color = ?", color)
	}

	if gender != "" {
		query = query.Where("gender = ?", gender)
	}

	if petType != "" {
		query = query.Where("pet_types.name = ?", petType)
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
	data := database.DB.Model(&models.Pet{}).Preload("PetType").Preload("Image")
	data = data.Preload("User").Preload("User.Image").Preload("User.Role")
	data = data.First(&pet, id)

	if data.Error != nil || data.RowsAffected == 0 {
		return models.Pet{}, fmt.Errorf("pet with id '%d' not found", id)
	}

	return pet, nil
}

func GetAllPets(offset, limit int, breed, color, gender, petType string) ([]models.Pet, error) {
	var pets []models.Pet
	query := database.DB.Model(&models.Pet{}).Joins("INNER JOIN pet_types ON pets.pet_type_id = pet_types.id")

	if breed != "" {
		query = query.Where("breed = ?", breed)
	}

	if color != "" {
		query = query.Where("color = ?", color)
	}

	if gender != "" {
		query = query.Where("gender = ?", gender)
	}

	if petType != "" {
		query = query.Where("pet_types.name = ?", petType)
	}

	data := query.Offset(offset).Limit(limit).Preload("User").Preload("Image")
	data = data.Preload("User.Image")

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
	operation := database.DB.Select("Image").Delete(&pet)

	if operation.Error != nil || operation.RowsAffected == 0 {
		return models.Pet{}, err
	}
	return pet, nil
}

func DeletePets(pets []models.Pet) ([]models.Pet, error) {
	if len(pets) == 0 {
		return []models.Pet{}, nil
	}
	operation := database.DB.Select("Image").Delete(&pets)
	if operation.Error != nil || operation.RowsAffected == 0 {
		return []models.Pet{}, operation.Error
	}
	return pets, nil
}
