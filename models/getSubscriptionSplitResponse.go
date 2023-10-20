package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

type GetSubscriptionSplitResponse struct {
	Enabled types.Optional[bool]               `json:"enabled"`
	Rules   types.Optional[[]GetSplitResponse] `json:"rules"`
}

func (g *GetSubscriptionSplitResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetSubscriptionSplitResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Enabled.IsValueSet() {
		structMap["enabled"] = g.Enabled.Value()
	}
	if g.Rules.IsValueSet() {
		structMap["rules"] = g.Rules.Value()
	}
	return structMap
}

func (g *GetSubscriptionSplitResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Enabled types.Optional[bool]               `json:"enabled"`
		Rules   types.Optional[[]GetSplitResponse] `json:"rules"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Enabled = temp.Enabled
	g.Rules = temp.Rules
	return nil
}
