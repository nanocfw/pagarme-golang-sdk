package models

import (
	"encoding/json"
)

type UpdateSubscriptionSplitRequest struct {
	Enabled bool                 `json:"enabled"`
	Rules   []CreateSplitRequest `json:"rules"`
}

func (u *UpdateSubscriptionSplitRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateSubscriptionSplitRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["enabled"] = u.Enabled
	structMap["rules"] = u.Rules
	return structMap
}

func (u *UpdateSubscriptionSplitRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Enabled bool                 `json:"enabled"`
		Rules   []CreateSplitRequest `json:"rules"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Enabled = temp.Enabled
	u.Rules = temp.Rules
	return nil
}
