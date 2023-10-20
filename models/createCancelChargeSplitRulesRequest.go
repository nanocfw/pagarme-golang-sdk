package models

import (
	"encoding/json"
)

// Creates a refund with split rules
type CreateCancelChargeSplitRulesRequest struct {
	Id     string `json:"id"`
	Amount int    `json:"Amount"`
	Type   string `json:"type"`
}

func (c *CreateCancelChargeSplitRulesRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCancelChargeSplitRulesRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["id"] = c.Id
	structMap["Amount"] = c.Amount
	structMap["type"] = c.Type
	return structMap
}

func (c *CreateCancelChargeSplitRulesRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id     string `json:"id"`
		Amount int    `json:"Amount"`
		Type   string `json:"type"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Id = temp.Id
	c.Amount = temp.Amount
	c.Type = temp.Type
	return nil
}
