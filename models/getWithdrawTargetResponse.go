package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

type GetWithdrawTargetResponse struct {
	TargetId types.Optional[string] `json:"target_id"`
	Type     types.Optional[string] `json:"type"`
}

func (g *GetWithdrawTargetResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetWithdrawTargetResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.TargetId.IsValueSet() {
		structMap["target_id"] = g.TargetId.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	return structMap
}

func (g *GetWithdrawTargetResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		TargetId types.Optional[string] `json:"target_id"`
		Type     types.Optional[string] `json:"type"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.TargetId = temp.TargetId
	g.Type = temp.Type
	return nil
}
