package mapper

import (
	"pet_api/src/dto/response"
	"pet_api/src/models"
)

func EventModelToResponse(event models.Event) response.EventResponse {
	return response.EventResponse{
		ID:              event.ID,
		Title:           event.Title,
		Description:     event.Description,
		ImageUrl:        event.ImageUrl,
		Date:            event.CreatedAt,
		ONG:             *OnlyUserModelToResponse(event.ONG),
		AllowVolunteers: event.AllowVolunteers,
		Participants:    OnlyUserModelsToResponse(event.Participants),
	}
}

func EventModelsToResponse(events []models.Event) []response.EventResponse {
	resp := make([]response.EventResponse, len(events))

	for i, v := range events {
		resp[i] = EventModelToResponse(v)
	}

	return resp
}
