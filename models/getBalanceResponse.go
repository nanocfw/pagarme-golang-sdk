package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Balance
type GetBalanceResponse struct {
	Currency           types.Optional[string]               `json:"currency"`
	AvailableAmount    types.Optional[int64]                `json:"available_amount"`
	Recipient          types.Optional[GetRecipientResponse] `json:"recipient"`
	TransferredAmount  types.Optional[int64]                `json:"transferred_amount"`
	WaitingFundsAmount types.Optional[int64]                `json:"waiting_funds_amount"`
}

func (g *GetBalanceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetBalanceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Currency.IsValueSet() {
		structMap["currency"] = g.Currency.Value()
	}
	if g.AvailableAmount.IsValueSet() {
		structMap["available_amount"] = g.AvailableAmount.Value()
	}
	if g.Recipient.IsValueSet() {
		structMap["recipient"] = g.Recipient.Value()
	}
	if g.TransferredAmount.IsValueSet() {
		structMap["transferred_amount"] = g.TransferredAmount.Value()
	}
	if g.WaitingFundsAmount.IsValueSet() {
		structMap["waiting_funds_amount"] = g.WaitingFundsAmount.Value()
	}
	return structMap
}

func (g *GetBalanceResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Currency           types.Optional[string]               `json:"currency"`
		AvailableAmount    types.Optional[int64]                `json:"available_amount"`
		Recipient          types.Optional[GetRecipientResponse] `json:"recipient"`
		TransferredAmount  types.Optional[int64]                `json:"transferred_amount"`
		WaitingFundsAmount types.Optional[int64]                `json:"waiting_funds_amount"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Currency = temp.Currency
	g.AvailableAmount = temp.AvailableAmount
	g.Recipient = temp.Recipient
	g.TransferredAmount = temp.TransferredAmount
	g.WaitingFundsAmount = temp.WaitingFundsAmount
	return nil
}
