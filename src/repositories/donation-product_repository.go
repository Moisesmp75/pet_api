package repositories

import (
	"fmt"
	"pet_api/src/database"
	"pet_api/src/models"
)

func CountDonationsProduct(ong_id uint64) int64 {
	var total_items int64
	query := database.DB.Model(&models.DonationProduct{})
	if ong_id != 0 {
		query = query.Where("ong_id = ?", ong_id)
	}
	if err := query.Count(&total_items).Error; err != nil {
		return 0
	}
	return int64(total_items)
}

func CreateDonationProduct(newDonationProduct models.DonationProduct) (models.DonationProduct, error) {
	if err := database.DB.Model(&models.DonationProduct{}).Create(&newDonationProduct).Preload("Products").Preload("User").Preload("Ong").Error; err != nil {
		return models.DonationProduct{}, err
	}

	return newDonationProduct, nil
}

func GetAllDonationsProduct(ong_id uint64, offset, limit int) ([]models.DonationProduct, error) {
	var donations []models.DonationProduct
	query := database.DB.Model(&models.DonationProduct{})
	if ong_id != 0 {
		query = query.Where("ong_id = ?", ong_id)
	}
	data := query.Offset(offset).Limit(limit)
	data = data.Preload("User").Preload("User.Role")
	data = data.Preload("Ong").Preload("Ong.Role")
	data = data.Preload("Products")
	data = data.Find(&donations)
	if data.Error != nil {
		return nil, data.Error
	}
	return donations, nil
}

func GetDonationProductById(id uint64) (models.DonationProduct, error) {
	var donation models.DonationProduct
	data := database.DB.Model(&models.DonationProduct{})
	data = data.Preload("User").Preload("User.Role")
	data = data.Preload("Ong").Preload("Ong.Role")
	data = data.Preload("Products")
	data = data.First(&donation, id)

	if data.Error != nil || data.RowsAffected == 0 {
		return models.DonationProduct{}, fmt.Errorf("donation with id '%d' not found", id)
	}

	return donation, nil
}
