package models

import (
	"encoding/json"
)

// Request for canceling a subscription
type CreateCancelSubscriptionRequest struct {
	CancelPendingInvoices bool `json:"cancel_pending_invoices"`
}

func (c *CreateCancelSubscriptionRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCancelSubscriptionRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["cancel_pending_invoices"] = c.CancelPendingInvoices
	return structMap
}

func (c *CreateCancelSubscriptionRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CancelPendingInvoices bool `json:"cancel_pending_invoices"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.CancelPendingInvoices = temp.CancelPendingInvoices
	return nil
}
