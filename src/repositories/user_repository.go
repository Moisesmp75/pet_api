package repositories

import (
	"fmt"
	"pet_api/src/database"
	"pet_api/src/models"
)

func CountUsers() int64 {
	var total_items int64
	database.DB.Model(&models.User{}).Count(&total_items)
	return total_items
}

func GetAllUsers(offset, limit int) ([]models.User, error) {
	var users []models.User
	data := database.DB.Model(&models.User{}).Offset(offset).Limit(limit).Preload("Role").Find(&users)
	if data.Error != nil {
		return nil, data.Error
	}
	return users, nil
}

func CreateUser(newUser models.User) (models.User, error) {
	if err := database.DB.Model(&models.User{}).Create(&newUser).Error; err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	data := database.DB.Model(&models.User{}).Where("email = ?", email).Preload("Role").First(&user)
	if data.Error != nil {
		return models.User{}, fmt.Errorf("user with email %s not found", email)
	}
	return user, nil
}

func GetUserById(id uint) (models.User, error) {
	var user models.User
	data := database.DB.Model(&models.User{}).Preload("Pets").Preload("Role").First(&user, id)
	if data.RowsAffected == 0 || data.Error != nil {
		return models.User{}, fmt.Errorf("user with ID %d not found", id)
	}
	return user, nil
}
