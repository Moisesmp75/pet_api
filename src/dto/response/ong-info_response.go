package response

type ONGInfoResponse struct {
	Description  string                `json:"description"`
	BankAccounts []BankAccountResponse `json:"bank_accounts"`
}
