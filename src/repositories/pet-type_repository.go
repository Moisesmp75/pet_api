package repositories

import (
	"fmt"
	"pet_api/src/database"
	"pet_api/src/models"
)

func GetPetTypeById(id uint64) (models.PetType, error) {
	var petType models.PetType
	data := database.DB.Model(&models.PetType{}).First(&petType, id)

	if data.Error != nil || data.RowsAffected == 0 {
		return models.PetType{}, fmt.Errorf("pet_type with id '%d' not found", id)
	}

	return petType, nil
}
