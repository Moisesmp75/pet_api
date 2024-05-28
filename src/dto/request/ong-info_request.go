package request

type ONGInfoRequest struct {
	Description string               `json:"description"`
	BankAccount []BankAccountRequest `json:"bank_accounts"`
}
