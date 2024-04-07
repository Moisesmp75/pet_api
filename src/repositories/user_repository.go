package repositories

import (
	"errors"
	"fmt"
	"pet_api/src/database"
	"pet_api/src/models"

	"gorm.io/gorm"
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

func GetUserByEmailOrPhone(identity string) (models.User, error) {
	var user models.User
	data := database.DB.Model(&models.User{}).Where("email = ? OR phone_number = ?", identity, identity).Preload("Role").First(&user)
	if data.Error != nil {
		if errors.Is(data.Error, gorm.ErrRecordNotFound) {
			return models.User{}, fmt.Errorf("user with email or phone number '%s' not found", identity)
		}
		return models.User{}, fmt.Errorf("failed to get user by email or phone number: %v", data.Error)
	}
	return user, nil
}

func GetUserById(id uint64) (models.User, error) {
	var user models.User
	data := database.DB.Model(&models.User{}).Preload("Pets").Preload("Role").First(&user, id)
	if data.RowsAffected == 0 || data.Error != nil {
		return models.User{}, fmt.Errorf("user with id '%d' not found", id)
	}
	return user, nil
}

func UpdateUser(user models.User) (models.User, error) {
	data := database.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(user)
	if data.RowsAffected == 0 || data.Error != nil {
		return models.User{}, data.Error
	}
	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	data := database.DB.Model(&models.User{}).Where("email = ?", email).Preload("Role").First(&user)
	if data.Error != nil {
		if errors.Is(data.Error, gorm.ErrRecordNotFound) {
			return models.User{}, fmt.Errorf("user with email '%s' not found", email)
		}
		return models.User{}, fmt.Errorf("failed to get user by email: %v", data.Error)
	}
	return user, nil
}
