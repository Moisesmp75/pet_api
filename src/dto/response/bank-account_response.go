package response

type BankAccountResponse struct {
	ID            uint64 `json:"id"`
	BankName      string `json:"bank_name"`
	AccountNumber string `json:"account_number"`
	CCI           string `json:"cci"`
}
