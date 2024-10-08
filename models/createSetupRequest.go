package models

import (
	"encoding/json"
)

// Request for creating a Setup for a subscription. The setup is an order that will be created at the subscription creation.
type CreateSetupRequest struct {
	Amount      int                  `json:"amount"`
	Description string               `json:"description"`
	Payment     CreatePaymentRequest `json:"payment"`
}

func (c *CreateSetupRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateSetupRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = c.Amount
	structMap["description"] = c.Description
	structMap["payment"] = c.Payment
	return structMap
}

func (c *CreateSetupRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount      int                  `json:"amount"`
		Description string               `json:"description"`
		Payment     CreatePaymentRequest `json:"payment"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Amount = temp.Amount
	c.Description = temp.Description
	c.Payment = temp.Payment
	return nil
}
