package repositories

import (
	"fmt"
	"pet_api/src/database"
	"pet_api/src/models"
)

func CreateOngInfo(ongInfo models.ONGInfo) (models.ONGInfo, error) {
	if err := database.DB.Model(&models.ONGInfo{}).Create(&ongInfo).Error; err != nil {
		return models.ONGInfo{}, err
	}
	return ongInfo, nil
}

func UpdateOngInfo(ongInfo models.ONGInfo) (models.ONGInfo, error) {
	data := database.DB.Model(&models.ONGInfo{}).Where("id = ?", ongInfo.ID).Updates(&ongInfo)
	if data.Error != nil || data.RowsAffected == 0 {
		return models.ONGInfo{}, fmt.Errorf("failed to update organization info: %v", data.Error)
	}
	return ongInfo, nil
}
