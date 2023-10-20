package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

type GetChargesSummaryResponse struct {
	Total types.Optional[int] `json:"total"`
}

func (g *GetChargesSummaryResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetChargesSummaryResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Total.IsValueSet() {
		structMap["total"] = g.Total.Value()
	}
	return structMap
}

func (g *GetChargesSummaryResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Total types.Optional[int] `json:"total"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Total = temp.Total
	return nil
}
