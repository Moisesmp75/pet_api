package repositories

import (
	"fmt"
	"pet_api/src/database"
	"pet_api/src/models"
)

func CountVisits() int64 {
	var total_items int64
	if err := database.DB.Model(&models.Visit{}).Count(&total_items).Error; err != nil {
		return 0
	}
	return total_items
}

func GetAllVisits(offset, limit int) ([]models.Visit, error) {
	var visits []models.Visit

	data := database.DB.Model(&models.Visit{})
	data = data.Offset(offset).Limit(limit)
	data = data.Preload("User").Preload("User.Role")
	data = data.Preload("Pet").Preload("Pet.User").Preload("Pet.User.Role").Preload("Pet.Behavior").Preload("Pet.Image")
	data = data.Find(&visits)

	if data.Error != nil {
		return []models.Visit{}, data.Error
	}

	return visits, nil
}

func GetVisitById(id uint64) (models.Visit, error) {
	var visit models.Visit
	data := database.DB.Model(&models.Visit{})
	data = data.Preload("User").Preload("User.Role")
	data = data.Preload("Pet").Preload("Pet.User").Preload("Pet.User.Role").Preload("Pet.Behavior").Preload("Pet.Image")
	data = data.First(&visit, id)

	if data.Error != nil || data.RowsAffected == 0 {
		return models.Visit{}, fmt.Errorf("visit with id '%d' not found", id)
	}

	return visit, nil
}

func CreateVisit(newVisit models.Visit) (models.Visit, error) {
	if err := database.DB.Model(&models.Visit{}).Create(&newVisit).Error; err != nil {
		return models.Visit{}, err
	}

	return newVisit, nil
}

func DeleteVisit(id uint64) (models.Visit, error) {
	visit, err := GetVisitById(id)
	if err != nil {
		return models.Visit{}, err
	}
	operation := database.DB.Model(&models.Visit{})
	operation = operation.Delete(&visit)

	if operation.Error != nil || operation.RowsAffected == 0 {
		return models.Visit{}, operation.Error
	}
	return visit, nil
}

func UpdateVisit(visit models.Visit) (models.Visit, error) {

	data := database.DB.Model(&models.Visit{}).Where("id = ?", visit.ID).Updates(visit)
	if data.RowsAffected == 0 || data.Error != nil {
		return models.Visit{}, data.Error
	}

	return visit, nil
}
