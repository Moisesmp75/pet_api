package mapper

import (
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/helpers"
	"pet_api/src/models"
)

func VisitRequestToModel(req request.VisitRequest) models.Visit {
	return models.Visit{
		PetID: req.PetID,
		Date:  helpers.ParseDateTime(req.Date),
		State: "Pending",
	}
}

func VisitModelToResponse(visit models.Visit) response.VisitResponse {
	return response.VisitResponse{
		ID:    visit.ID,
		Pet:   PetModelToResponse(visit.Pet),
		User:  *OnlyUserModelToResponse(visit.User),
		Date:  visit.Date,
		State: visit.State,
	}
}

func VisitModelsToResponse(visits []models.Visit) []response.VisitResponse {
	resp := make([]response.VisitResponse, len(visits))

	for i, v := range visits {
		resp[i] = VisitModelToResponse(v)
	}

	return resp
}

func UpdateVisitRequestToModel(req request.UpdateVisitRequest, visit models.Visit) models.Visit {
	if req.State != "" {
		visit.State = req.State
	}
	return visit
}
