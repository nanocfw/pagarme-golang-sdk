package models

import (
	"encoding/json"
)

// Token data
type CreateTokenRequest struct {
	Type string                 `json:"type"`
	Card CreateCardTokenRequest `json:"card"`
}

func (c *CreateTokenRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateTokenRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["type"] = c.Type
	structMap["card"] = c.Card
	return structMap
}

func (c *CreateTokenRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Type string                 `json:"type"`
		Card CreateCardTokenRequest `json:"card"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Type = temp.Type
	c.Card = temp.Card
	return nil
}
