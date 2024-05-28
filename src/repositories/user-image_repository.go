package repositories

import (
	"pet_api/src/database"
	"pet_api/src/models"
)

func UpdateUserImage(userImage models.UserImage) (models.UserImage, error) {
	data := database.DB.Model(&models.UserImage{}).Where("id = ?", userImage.ID).Updates(userImage)
	if data.RowsAffected == 0 || data.Error != nil {
		return models.UserImage{}, data.Error
	}
	return userImage, nil
}
