package models

import (
	"encoding/json"
)

type CreateConfirmPaymentRequest struct {
	Description string `json:"description"`
	Amount      *int   `json:"Amount,omitempty"`
	Code        string `json:"Code"`
}

func (c *CreateConfirmPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateConfirmPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["description"] = c.Description
	if c.Amount != nil {
		structMap["Amount"] = c.Amount
	}
	structMap["Code"] = c.Code
	return structMap
}

func (c *CreateConfirmPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Description string `json:"description"`
		Amount      *int   `json:"Amount,omitempty"`
		Code        string `json:"Code"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Description = temp.Description
	c.Amount = temp.Amount
	c.Code = temp.Code
	return nil
}
