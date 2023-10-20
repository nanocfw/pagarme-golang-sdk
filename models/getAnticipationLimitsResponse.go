package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Anticipation limits
type GetAnticipationLimitsResponse struct {
	Max types.Optional[GetAnticipationLimitResponse] `json:"max"`
	Min types.Optional[GetAnticipationLimitResponse] `json:"min"`
}

func (g *GetAnticipationLimitsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetAnticipationLimitsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Max.IsValueSet() {
		structMap["max"] = g.Max.Value()
	}
	if g.Min.IsValueSet() {
		structMap["min"] = g.Min.Value()
	}
	return structMap
}

func (g *GetAnticipationLimitsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Max types.Optional[GetAnticipationLimitResponse] `json:"max"`
		Min types.Optional[GetAnticipationLimitResponse] `json:"min"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Max = temp.Max
	g.Min = temp.Min
	return nil
}
