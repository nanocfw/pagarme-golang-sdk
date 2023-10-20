package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Information about fines and interest on the "boleto" used from payment
type CreateSubscriptionBoletoRequest struct {
	Interest            *CreateInterestRequest `json:"interest,omitempty"`
	Fine                *CreateFineRequest     `json:"fine,omitempty"`
	MaxDaysToPayPastDue types.Optional[int]    `json:"max_days_to_pay_past_due"`
}

func (c *CreateSubscriptionBoletoRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateSubscriptionBoletoRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Interest != nil {
		structMap["interest"] = c.Interest
	}
	if c.Fine != nil {
		structMap["fine"] = c.Fine
	}
	if c.MaxDaysToPayPastDue.IsValueSet() {
		structMap["max_days_to_pay_past_due"] = c.MaxDaysToPayPastDue.Value()
	}
	return structMap
}

func (c *CreateSubscriptionBoletoRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Interest            *CreateInterestRequest `json:"interest,omitempty"`
		Fine                *CreateFineRequest     `json:"fine,omitempty"`
		MaxDaysToPayPastDue types.Optional[int]    `json:"max_days_to_pay_past_due"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Interest = temp.Interest
	c.Fine = temp.Fine
	c.MaxDaysToPayPastDue = temp.MaxDaysToPayPastDue
	return nil
}
