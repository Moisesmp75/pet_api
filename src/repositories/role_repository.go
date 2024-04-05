package repositories

import (
	"pet_api/src/database"
	"pet_api/src/models"
)

func GetRoleById(id uint) (models.Role, error) {
	var role models.Role
	data := database.DB.Model(&models.Role{}).First(&role, id)

	if data.Error != nil || data.RowsAffected == 0 {
		return models.Role{}, data.Error
	}
	return role, nil
}