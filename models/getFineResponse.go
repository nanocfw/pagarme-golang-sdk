package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Fine Response
type GetFineResponse struct {
	Days   types.Optional[int]    `json:"days"`
	Type   types.Optional[string] `json:"type"`
	Amount types.Optional[int]    `json:"amount"`
}

func (g *GetFineResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetFineResponse) toMap() map[string]any {
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

func (g *GetFineResponse) UnmarshalJSON(input []byte) error {
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
