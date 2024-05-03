package mapper

import (
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/models"
)

func EventRequestToModel(req request.EventRequest) models.Event {
	return models.Event{
		Title:           req.Title,
		Description:     req.Description,
		AllowVolunteers: req.AllowVolunteers,
	}
}

func EventModelToResponse(event models.Event) response.EventResponse {
	return response.EventResponse{
		ID:              event.ID,
		Title:           event.Title,
		Description:     event.Description,
		ImageUrl:        event.ImageUrl,
		PublicationDate: event.CreatedAt,
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
