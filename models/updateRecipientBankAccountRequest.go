package models

import (
	"encoding/json"
)

// Updates the default bank account for a recipient
type UpdateRecipientBankAccountRequest struct {
	BankAccount CreateBankAccountRequest `json:"bank_account"`
	PaymentMode string                   `json:"payment_mode"`
}

func (u *UpdateRecipientBankAccountRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateRecipientBankAccountRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["bank_account"] = u.BankAccount
	structMap["payment_mode"] = u.PaymentMode
	return structMap
}

func (u *UpdateRecipientBankAccountRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		BankAccount CreateBankAccountRequest `json:"bank_account"`
		PaymentMode string                   `json:"payment_mode"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.BankAccount = temp.BankAccount
	u.PaymentMode = temp.PaymentMode
	return nil
}
