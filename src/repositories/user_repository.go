package repositories

import (
	"pet_api/src/database"
	"pet_api/src/models"
)

func CountUsers() int64 {
	var total_items int64
	database.DB.Model(&models.User{}).Count(&total_items)
	return total_items
}

func GetAllUsers(offset, limit int) (*[]models.User, error) {
	var users []models.User
	data := database.DB.Offset(offset).Limit(limit).Find(&users)
	if data.Error != nil {
		return nil, data.Error
	}
	return &users, nil
}

func CreateUser(newUser models.User) (*models.User, error) {
	if err := database.DB.Create(&newUser).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	data := database.DB.Where("email = ?", email).First(&user)
	if data.Error != nil {
		return nil, data.Error
	}
	return &user, nil
}

func GetUserById(id uint) (*models.User, error) {
	var user models.User
	data := database.DB.Find(&user, id)
	if data.Error != nil || data.RowsAffected == 0 {
		return nil, data.Error
	}
	return &user, nil
}
