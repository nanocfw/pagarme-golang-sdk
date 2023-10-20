package models

import (
	"encoding/json"
)

// Fine Request
type CreateFineRequest struct {
	Days   int    `json:"days"`
	Type   string `json:"type"`
	Amount int    `json:"amount"`
}

func (c *CreateFineRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateFineRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["days"] = c.Days
	structMap["type"] = c.Type
	structMap["amount"] = c.Amount
	return structMap
}

func (c *CreateFineRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Days   int    `json:"days"`
		Type   string `json:"type"`
		Amount int    `json:"amount"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Days = temp.Days
	c.Type = temp.Type
	c.Amount = temp.Amount
	return nil
}
