package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

type GetCheckoutBoletoPaymentResponse struct {
	DueAt        types.Optional[time.Time] `json:"due_at"`
	Instructions types.Optional[string]    `json:"instructions"`
}

func (g *GetCheckoutBoletoPaymentResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetCheckoutBoletoPaymentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.DueAt.IsValueSet() {
		var DueAtVal *string = nil
		if g.DueAt.Value() != nil {
			val := g.DueAt.Value().Format(time.RFC3339)
			DueAtVal = &val
		}
		structMap["due_at"] = DueAtVal
	}
	if g.Instructions.IsValueSet() {
		structMap["instructions"] = g.Instructions.Value()
	}
	return structMap
}

func (g *GetCheckoutBoletoPaymentResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		DueAt        types.Optional[string] `json:"due_at"`
		Instructions types.Optional[string] `json:"instructions"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.DueAt.ShouldSetValue(temp.DueAt.IsValueSet())
	if temp.DueAt.Value() != nil {
		DueAtVal, err := time.Parse(time.RFC3339, (*temp.DueAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
		}
		g.DueAt.SetValue(&DueAtVal)
	}
	g.Instructions = temp.Instructions
	return nil
}
