package models

import (
	"encoding/json"
)

// Request for creating a transfer
type CreateTransferRequest struct {
	Amount   int               `json:"amount"`
	Metadata map[string]string `json:"metadata"`
}

func (c *CreateTransferRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateTransferRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = c.Amount
	structMap["metadata"] = c.Metadata
	return structMap
}

func (c *CreateTransferRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount   int               `json:"amount"`
		Metadata map[string]string `json:"metadata"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Amount = temp.Amount
	c.Metadata = temp.Metadata
	return nil
}
