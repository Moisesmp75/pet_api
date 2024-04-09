package repositories

import (
	"fmt"
	"pet_api/src/database"
	"pet_api/src/models"
)

func GetRoleById(id uint64) (models.Role, error) {
	var role models.Role
	data := database.DB.Model(&models.Role{}).First(&role, id)

	if data.Error != nil || data.RowsAffected == 0 {
		return models.Role{}, fmt.Errorf("rol with id '%d' not found", id)
	}
	return role, nil
}
