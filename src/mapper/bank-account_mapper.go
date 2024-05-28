package mapper

import (
	"pet_api/src/dto/response"
	"pet_api/src/models"
)

func BankAccountModelToResponse(bank_account models.BankAccount) response.BankAccountResponse {
	return response.BankAccountResponse{
		ID:            bank_account.ID,
		BankName:      bank_account.BankName,
		AccountNumber: bank_account.AccountNumber,
		CCI:           bank_account.CCI,
	}
}

func BankAccountModelsToResponse(bank_accounts []models.BankAccount) []response.BankAccountResponse {
	resp := make([]response.BankAccountResponse, len(bank_accounts))

	for i, v := range bank_accounts {
		resp[i] = BankAccountModelToResponse(v)
	}
	return resp
}
