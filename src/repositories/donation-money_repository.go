package repositories

import (
	"fmt"
	"pet_api/src/database"
	"pet_api/src/models"
)

func CountDonationsMoney(ongID uint64) int64 {
	var totalItems int64
	query := database.DB.Model(&models.DonationMoney{})
	if ongID != 0 {
		query = query.Where("ong_id = ?", ongID)
	}
	if err := query.Count(&totalItems).Error; err != nil {
		return 0
	}
	return totalItems
}

func CreateDonationMoney(newDonationMoney models.DonationMoney) (models.DonationMoney, error) {
	if err := database.DB.Model(&models.DonationMoney{}).Create(&newDonationMoney).Error; err != nil {
		return models.DonationMoney{}, err
	}
	return newDonationMoney, nil
}

func GetAllDonationsMoney(ongID uint64, offset, limit int) ([]models.DonationMoney, error) {
	var donations []models.DonationMoney
	query := database.DB.Model(&models.DonationMoney{})
	if ongID != 0 {
		query = query.Where("ong_id = ?", ongID)
	}
	data := query.Offset(offset).Limit(limit).Preload("User").Preload("User.Role").Preload("Ong").Preload("Ong.Role").Find(&donations)
	if data.Error != nil {
		return nil, data.Error
	}
	return donations, nil
}

func GetDonationMoneyByID(id uint64) (models.DonationMoney, error) {
	var donation models.DonationMoney
	data := database.DB.Model(&models.DonationMoney{}).Preload("User").Preload("User.Role").Preload("Ong").Preload("Ong.Role").First(&donation, id)
	if data.Error != nil || data.RowsAffected == 0 {
		return models.DonationMoney{}, fmt.Errorf("donation with ID '%d' not found", id)
	}
	return donation, nil
}

func UpdateDonationMoney(donation models.DonationMoney) (models.DonationMoney, error) {
	data := database.DB.Model(&models.DonationMoney{}).Where("id = ?", donation.ID).Updates(donation)
	if data.RowsAffected == 0 || data.Error != nil {
		return models.DonationMoney{}, data.Error
	}
	return donation, nil
}

func DeleteDonationMoney(id uint64) (models.DonationMoney, error) {
	donation, err := GetDonationMoneyByID(id)
	if err != nil {
		return models.DonationMoney{}, err
	}
	operation := database.DB.Model(&models.DonationMoney{}).Delete(&donation)
	if operation.Error != nil || operation.RowsAffected == 0 {
		return models.DonationMoney{}, operation.Error
	}
	return donation, nil
}
