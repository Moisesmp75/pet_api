package mapper

import (
	"pet_api/src/dto/response"
	"pet_api/src/models"
)

func ONGInfoModelToResponse(ONGInfo models.ONGInfo) response.ONGInfoResponse {
	return response.ONGInfoResponse{
		Description:  ONGInfo.Description,
		BankAccounts: BankAccountModelsToResponse(ONGInfo.BankAccounts),
	}
}
