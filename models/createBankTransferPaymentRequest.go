package models

import (
	"encoding/json"
)

// Request for creating a bank transfer payment
type CreateBankTransferPaymentRequest struct {
	Bank    string `json:"bank"`
	Retries int    `json:"retries"`
}

func (c *CreateBankTransferPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateBankTransferPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["bank"] = c.Bank
	structMap["retries"] = c.Retries
	return structMap
}

func (c *CreateBankTransferPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Bank    string `json:"bank"`
		Retries int    `json:"retries"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Bank = temp.Bank
	c.Retries = temp.Retries
	return nil
}
