package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

type GetCheckoutDebitCardPaymentResponse struct {
	StatementDescriptor types.Optional[string]                           `json:"statement_descriptor"`
	Authentication      types.Optional[GetPaymentAuthenticationResponse] `json:"authentication"`
}

func (g *GetCheckoutDebitCardPaymentResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetCheckoutDebitCardPaymentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.StatementDescriptor.IsValueSet() {
		structMap["statement_descriptor"] = g.StatementDescriptor.Value()
	}
	if g.Authentication.IsValueSet() {
		structMap["authentication"] = g.Authentication.Value()
	}
	return structMap
}

func (g *GetCheckoutDebitCardPaymentResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StatementDescriptor types.Optional[string]                           `json:"statement_descriptor"`
		Authentication      types.Optional[GetPaymentAuthenticationResponse] `json:"authentication"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.StatementDescriptor = temp.StatementDescriptor
	g.Authentication = temp.Authentication
	return nil
}
