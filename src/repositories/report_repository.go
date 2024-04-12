package repositories

import (
	"fmt"
	"pet_api/src/database"
	"pet_api/src/models"
)

func CountReports() int64 {
	var total_items int64
	if err := database.DB.Model(&models.Report{}).Count(&total_items).Error; err != nil {
		return 0
	}
	return total_items
}

func GetAllReports(offset, limit int) ([]models.Report, error) {
	var reports []models.Report

	data := database.DB.Model(&models.Report{})
	data = data.Offset(offset).Limit(limit)
	data = data.Preload("ReporterUser").Preload("ReporterUser.Role")
	data = data.Preload("ReportedUser").Preload("ReportedUser.Role")
	data = data.Find(&reports)

	if data.Error != nil {
		return []models.Report{}, data.Error
	}

	return reports, nil
}

func GetReportById(id uint64) (models.Report, error) {
	var report models.Report
	data := database.DB.Model(&models.Report{})
	data = data.Preload("ReporterUser").Preload("ReporterUser.Role")
	data = data.Preload("ReportedUser").Preload("ReportedUser.Role")
	data = data.First(&report, id)

	if data.Error != nil || data.RowsAffected == 0 {
		return models.Report{}, fmt.Errorf("report with id '%v' not found", id)
	}

	return report, nil
}

func CreateReport(newReport models.Report) (models.Report, error) {
	if err := database.DB.Model(&models.Report{}).Create(newReport).Error; err != nil {
		return models.Report{}, err
	}
	return newReport, nil
}
