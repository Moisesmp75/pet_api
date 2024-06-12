package mapper

import (
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/models"
)

func DonationMoneyRequestToModel(req request.DonationMoneyRequest) models.DonationMoney {
	return models.DonationMoney{
		OngID:    req.OngID,
		Amount:   req.Amount,
		Received: false,
	}
}

func DonationMoneyModelToResponse(donation models.DonationMoney) response.DonationMoneyResponse {
	return response.DonationMoneyResponse{
		ID:           donation.ID,
		User:         *OnlyUserModelToResponse(donation.User),
		Ong:          *OnlyUserModelToResponse(donation.Ong),
		Amount:       donation.Amount,
		DonationDate: donation.CreatedAt,
		Received:     donation.Received,
	}
}

func DonationMoneyModelsToResponse(donations []models.DonationMoney) []response.DonationMoneyResponse {
	resp := make([]response.DonationMoneyResponse, len(donations))

	for i, v := range donations {
		resp[i] = DonationMoneyModelToResponse(v)
	}
	return resp
}

func UpdateDonationMoneyRequestToModel(req request.UpdateDonationMoneyRequest, donation models.DonationMoney) models.DonationMoney {
	donation.Received = req.Received
	return donation
}
