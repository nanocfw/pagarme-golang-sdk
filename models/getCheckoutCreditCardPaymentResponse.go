package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

type GetCheckoutCreditCardPaymentResponse struct {
	StatementDescriptor types.Optional[string]                                      `json:"statementDescriptor"`
	Installments        types.Optional[[]GetCheckoutCardInstallmentOptionsResponse] `json:"installments"`
	Authentication      types.Optional[GetPaymentAuthenticationResponse]            `json:"authentication"`
}

func (g *GetCheckoutCreditCardPaymentResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetCheckoutCreditCardPaymentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.StatementDescriptor.IsValueSet() {
		structMap["statementDescriptor"] = g.StatementDescriptor.Value()
	}
	if g.Installments.IsValueSet() {
		structMap["installments"] = g.Installments.Value()
	}
	if g.Authentication.IsValueSet() {
		structMap["authentication"] = g.Authentication.Value()
	}
	return structMap
}

func (g *GetCheckoutCreditCardPaymentResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StatementDescriptor types.Optional[string]                                      `json:"statementDescriptor"`
		Installments        types.Optional[[]GetCheckoutCardInstallmentOptionsResponse] `json:"installments"`
		Authentication      types.Optional[GetPaymentAuthenticationResponse]            `json:"authentication"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.StatementDescriptor = temp.StatementDescriptor
	g.Installments = temp.Installments
	g.Authentication = temp.Authentication
	return nil
}
