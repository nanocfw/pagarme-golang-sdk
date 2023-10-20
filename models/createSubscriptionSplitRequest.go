package models

import (
	"encoding/json"
)

type CreateSubscriptionSplitRequest struct {
	Enabled bool                 `json:"enabled"`
	Rules   []CreateSplitRequest `json:"rules"`
}

func (c *CreateSubscriptionSplitRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateSubscriptionSplitRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["enabled"] = c.Enabled
	structMap["rules"] = c.Rules
	return structMap
}

func (c *CreateSubscriptionSplitRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Enabled bool                 `json:"enabled"`
		Rules   []CreateSplitRequest `json:"rules"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Enabled = temp.Enabled
	c.Rules = temp.Rules
	return nil
}
