package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Anticipation limit
type GetAnticipationLimitResponse struct {
	Amount          types.Optional[int] `json:"amount"`
	AnticipationFee types.Optional[int] `json:"anticipation_fee"`
}

func (g *GetAnticipationLimitResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetAnticipationLimitResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.AnticipationFee.IsValueSet() {
		structMap["anticipation_fee"] = g.AnticipationFee.Value()
	}
	return structMap
}

func (g *GetAnticipationLimitResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount          types.Optional[int] `json:"amount"`
		AnticipationFee types.Optional[int] `json:"anticipation_fee"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Amount = temp.Amount
	g.AnticipationFee = temp.AnticipationFee
	return nil
}
