package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a boleto
type GetSubscriptionBoletoResponse struct {
	Interest            types.Optional[GetInterestResponse] `json:"interest"`
	Fine                types.Optional[GetFineResponse]     `json:"fine"`
	MaxDaysToPayPastDue types.Optional[int]                 `json:"max_days_to_pay_past_due"`
}

func (g *GetSubscriptionBoletoResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetSubscriptionBoletoResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Interest.IsValueSet() {
		structMap["interest"] = g.Interest.Value()
	}
	if g.Fine.IsValueSet() {
		structMap["fine"] = g.Fine.Value()
	}
	if g.MaxDaysToPayPastDue.IsValueSet() {
		structMap["max_days_to_pay_past_due"] = g.MaxDaysToPayPastDue.Value()
	}
	return structMap
}

func (g *GetSubscriptionBoletoResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Interest            types.Optional[GetInterestResponse] `json:"interest"`
		Fine                types.Optional[GetFineResponse]     `json:"fine"`
		MaxDaysToPayPastDue types.Optional[int]                 `json:"max_days_to_pay_past_due"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Interest = temp.Interest
	g.Fine = temp.Fine
	g.MaxDaysToPayPastDue = temp.MaxDaysToPayPastDue
	return nil
}
