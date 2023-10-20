package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Interest Response
type GetInterestResponse struct {
	Days   types.Optional[int]    `json:"days"`
	Type   types.Optional[string] `json:"type"`
	Amount types.Optional[int]    `json:"amount"`
}

func (g *GetInterestResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetInterestResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Days.IsValueSet() {
		structMap["days"] = g.Days.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	return structMap
}

func (g *GetInterestResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Days   types.Optional[int]    `json:"days"`
		Type   types.Optional[string] `json:"type"`
		Amount types.Optional[int]    `json:"amount"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Days = temp.Days
	g.Type = temp.Type
	g.Amount = temp.Amount
	return nil
}
