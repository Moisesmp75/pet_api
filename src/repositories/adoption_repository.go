package repositories

import (
	"fmt"
	"pet_api/src/database"
	"pet_api/src/models"
)

func CountAdoptions() int64 {
	var total_items int64
	if err := database.DB.Model(&models.Adoption{}).Count(&total_items).Where("deleted_at = ?", nil).Error; err != nil {
		return 0
	}
	return int64(total_items)
}

func GetAllAdoptions(offset, limit int) ([]models.Adoption, error) {
	var adoptions []models.Adoption

	data := database.DB.Model(&models.Adoption{})
	data = data.Offset(offset).Limit(limit)
	data = data.Preload("Pet").Preload("Pet.User").Preload("Pet.User.Role").Preload("Pet.Behavior").Preload("Pet.Image")
	data = data.Preload("User").Preload("User.Role").Preload("User.Image")
	data = data.Find(&adoptions)

	if data.Error != nil {
		return []models.Adoption{}, nil
	}

	return adoptions, nil
}

func GetAdoptionById(id uint64) (models.Adoption, error) {
	var adoption models.Adoption

	data := database.DB.Model(&models.Adoption{})
	data = data.Preload("Pet").Preload("Pet.User").Preload("Pet.User.Role").Preload("Pet.Behavior").Preload("Pet.Image")
	data = data.Preload("User").Preload("User.Role").Preload("User.Image")
	data = data.First(&adoption, id)

	if data.Error != nil || data.RowsAffected == 0 {
		return models.Adoption{}, fmt.Errorf("adoption with id '%v' not found", id)
	}

	return adoption, nil
}

func CreateAdoption(newAdoption models.Adoption) (models.Adoption, error) {
	if err := database.DB.Model(&models.Adoption{}).Create(&newAdoption).Error; err != nil {
		return models.Adoption{}, err
	}
	return newAdoption, nil
}

func DeleteAdoption(id uint64) (models.Adoption, error) {
	adoption, err := GetAdoptionById(id)
	if err != nil {
		return models.Adoption{}, err
	}
	operation := database.DB.Model(&models.Adoption{})
	operation = operation.Delete(&adoption)

	if operation.Error != nil || operation.RowsAffected == 0 {
		return models.Adoption{}, operation.Error
	}
	return adoption, nil
}

func UpdateAdoption(adoption models.Adoption) (models.Adoption, error) {
	data := database.DB.Model(&models.Adoption{}).Where("id = ?", adoption.ID).Updates(adoption)
	if data.RowsAffected == 0 || data.Error != nil {
		return models.Adoption{}, data.Error
	}

	return adoption, nil
}
