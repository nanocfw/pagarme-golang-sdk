package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

type GetSplitOptionsResponse struct {
	Liable              types.Optional[bool]   `json:"liable"`
	ChargeProcessingFee types.Optional[bool]   `json:"charge_processing_fee"`
	ChargeRemainderFee  types.Optional[string] `json:"charge_remainder_fee"`
}

func (g *GetSplitOptionsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetSplitOptionsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Liable.IsValueSet() {
		structMap["liable"] = g.Liable.Value()
	}
	if g.ChargeProcessingFee.IsValueSet() {
		structMap["charge_processing_fee"] = g.ChargeProcessingFee.Value()
	}
	if g.ChargeRemainderFee.IsValueSet() {
		structMap["charge_remainder_fee"] = g.ChargeRemainderFee.Value()
	}
	return structMap
}

func (g *GetSplitOptionsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Liable              types.Optional[bool]   `json:"liable"`
		ChargeProcessingFee types.Optional[bool]   `json:"charge_processing_fee"`
		ChargeRemainderFee  types.Optional[string] `json:"charge_remainder_fee"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Liable = temp.Liable
	g.ChargeProcessingFee = temp.ChargeProcessingFee
	g.ChargeRemainderFee = temp.ChargeRemainderFee
	return nil
}
