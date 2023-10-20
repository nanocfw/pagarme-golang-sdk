package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Bank transfer checkout response
type GetCheckoutBankTransferPaymentResponse struct {
	Bank types.Optional[[]string] `json:"bank"`
}

func (g *GetCheckoutBankTransferPaymentResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetCheckoutBankTransferPaymentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Bank.IsValueSet() {
		structMap["bank"] = g.Bank.Value()
	}
	return structMap
}

func (g *GetCheckoutBankTransferPaymentResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Bank types.Optional[[]string] `json:"bank"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Bank = temp.Bank
	return nil
}
