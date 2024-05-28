package request

type BankAccountRequest struct {
	BankName      string `json:"bank_name" validate:"required,gt=0"`
	AccountNumber string `json:"account_number" validate:"required,gt=0"`
	CCI           string `json:"cci" validate:"required,gt=0"`
}
