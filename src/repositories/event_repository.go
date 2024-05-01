package repositories

import (
	"pet_api/src/database"
	"pet_api/src/models"
)

func CountEvents() int64 {
	var total_items int64
	if err := database.DB.Model(&models.Event{}).Count(&total_items).Error; err != nil {
		return 0
	}
	return int64(total_items)
}

func GetAllEvents(offset, limit int) ([]models.Event, error) {
	var events []models.Event

	data := database.DB.Model(&models.Event{})
	data = data.Offset(offset).Limit(limit)
	data = data.Preload("ONG").Preload("ONG.Role")
	data = data.Preload("Participants")
	data = data.Find(&events)

	if data.Error != nil {
		return []models.Event{}, data.Error
	}

	return events, nil
}

func CreateEvent(newEvent models.Event) (models.Event, error) {
	if err := database.DB.Model(&models.Event{}).Create(&newEvent).Error; err != nil {
		return models.Event{}, err
	}
	return newEvent, nil
}

func GetEventById(id uint64) (models.Event, error) {
	var event models.Event
	data := database.DB.Model(&models.Event{})
	data = data.Preload("ONG").Preload("ONG.Role")
	data = data.Preload("Participants")
	data = data.First(&event, id)

	if data.Error != nil || data.RowsAffected != 0 {
		return models.Event{}, data.Error
	}

	return event, nil
}

func AddParticipant(eventID uint64, userID uint64) (bool, error) {

	event, err := GetEventById(eventID)
	if err != nil {
		return false, err
	}

	user, err := GetUserById(userID)
	if err != nil {
		return false, err
	}

	if err := database.DB.Model(&event).Association("Participants").Append(&user); err != nil {
		return false, err
	}
	return true, nil
}

func DeleteEvent(id uint64) (models.Event, error) {
	event, err := GetEventById(id)
	if err != nil {
		return models.Event{}, err
	}

	operation := database.DB.Model(&models.Event{})
	operation = operation.Delete(&event)

	if operation.Error != nil || operation.RowsAffected == 0 {
		return models.Event{}, operation.Error
	}
	return event, nil
}
