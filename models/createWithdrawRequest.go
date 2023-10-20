package models

import (
	"encoding/json"
)

type CreateWithdrawRequest struct {
	Amount   int                `json:"amount"`
	Metadata map[string]*string `json:"metadata,omitempty"`
}

func (c *CreateWithdrawRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateWithdrawRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = c.Amount
	if c.Metadata != nil {
		structMap["metadata"] = c.Metadata
	}
	return structMap
}

func (c *CreateWithdrawRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount   int                `json:"amount"`
		Metadata map[string]*string `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Amount = temp.Amount
	c.Metadata = temp.Metadata
	return nil
}
