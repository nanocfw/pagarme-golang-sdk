package models

import (
	"encoding/json"
)

// Split
type CreateSplitRequest struct {
	Type        string                     `json:"type"`
	Amount      int                        `json:"amount"`
	RecipientId string                     `json:"recipient_id"`
	Options     *CreateSplitOptionsRequest `json:"options,omitempty"`
	SplitRuleId *string                    `json:"split_rule_id,omitempty"`
}

func (c *CreateSplitRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateSplitRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["type"] = c.Type
	structMap["amount"] = c.Amount
	structMap["recipient_id"] = c.RecipientId
	if c.Options != nil {
		structMap["options"] = c.Options
	}
	if c.SplitRuleId != nil {
		structMap["split_rule_id"] = c.SplitRuleId
	}
	return structMap
}

func (c *CreateSplitRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Type        string                     `json:"type"`
		Amount      int                        `json:"amount"`
		RecipientId string                     `json:"recipient_id"`
		Options     *CreateSplitOptionsRequest `json:"options,omitempty"`
		SplitRuleId *string                    `json:"split_rule_id,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Type = temp.Type
	c.Amount = temp.Amount
	c.RecipientId = temp.RecipientId
	c.Options = temp.Options
	c.SplitRuleId = temp.SplitRuleId
	return nil
}
