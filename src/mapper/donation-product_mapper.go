package mapper

import (
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/models"
)

func DonationProductRequestToModel(req request.DonationProductRequest) models.DonationProduct {
	return models.DonationProduct{
		OngID:    req.OngId,
		Products: ProductRequestsToModels(req.Products),
		Received: false,
	}
}

func DonationProductModelToResponse(donation models.DonationProduct) response.DonationProductResponse {
	return response.DonationProductResponse{
		ID:           donation.ID,
		User:         *OnlyUserModelToResponse(donation.User),
		Ong:          *OnlyUserModelToResponse(donation.Ong),
		Products:     ProductModelsToResponse(donation.Products),
		DonationDate: donation.CreatedAt,
		Received:     donation.Received,
	}
}

func DonationProductModelsToResponse(donations []models.DonationProduct) []response.DonationProductResponse {
	resp := make([]response.DonationProductResponse, len(donations))

	for i, v := range donations {
		resp[i] = DonationProductModelToResponse(v)
	}
	return resp
}

func UpdateDonationRequestToModel(req request.UpdateDonationRequest, donation models.DonationProduct) models.DonationProduct {
	donation.Received = req.Received
	return donation
}
