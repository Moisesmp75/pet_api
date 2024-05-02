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
	if err := database.DB.Model(&models.User{}).Count(&total_items).Error; err != nil {
		return 0
	}
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
	tx := database.DB.Begin()

	if err := tx.Model(&models.User{}).Create(&newUser).Preload("Role").Error; err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	newOngInfo := models.ONGInfo{UserID: newUser.ID}

	if err := tx.Model(&models.ONGInfo{}).Create(&newOngInfo).Preload("BankAccounts").Error; err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	tx.Commit()
	// if err := database.DB.Model(&models.User{}).Create(&newUser).Preload("Role").Error; err != nil {
	// 	return models.User{}, err
	// }
	return newUser, nil
}

func GetUserByEmailOrPhone(identity string) (models.User, error) {
	var user models.User
	data := database.DB.Model(&models.User{}).Where("email = ? OR phone_number = ?", identity, identity)
	data = data.Preload("Role")
	data = data.Preload("Pets").Preload("Pets.PetType").Preload("Pets.Images")
	data = data.Preload("ONGInfo").Preload("ONGInfo.BankAccounts")
	data = data.First(&user)
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
	data := database.DB.Model(&models.User{})
	data = data.Preload("Pets").Preload("Pets.PetType").Preload("Pets.Images")
	data = data.Preload("Role")
	data = data.Preload("ONGInfo").Preload("ONGInfo.BankAccounts")
	data = data.First(&user, id)
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
	data := database.DB.Model(&models.User{}).Where("email = ?", email)
	data = data.Preload("Role")
	data = data.Preload("Pets").Preload("Pets.PetType").Preload("Pets.Images")
	data = data.Preload("ONGInfo").Preload("ONGInfo.BankAccounts")
	data = data.First(&user)
	if data.Error != nil {
		if errors.Is(data.Error, gorm.ErrRecordNotFound) {
			return models.User{}, fmt.Errorf("user with email '%s' not found", email)
		}
		return models.User{}, fmt.Errorf("failed to get user by email: %v", data.Error)
	}
	return user, nil
}

func DeleteUser(id uint64) (models.User, error) {
	user, err := GetUserById(id)

	if err != nil {
		return models.User{}, err
	}

	if _, err := DeletePets(user.Pets); err != nil {
		return models.User{}, err
	}

	operation := database.DB.Select("Pets").Select("Pets.Images")
	operation = operation.Select("ONGInfo").Select("ONGInfo.BankAccounts")
	operation = operation.Delete(&user)

	if operation.Error != nil || operation.RowsAffected == 0 {
		return models.User{}, err
	}
	return user, nil
}
