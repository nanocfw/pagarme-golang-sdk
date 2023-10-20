package models

import (
	"encoding/json"
	"log"
	"time"
)

// Request for updating the due date from a subscription
type UpdateSubscriptionBillingDateRequest struct {
	NextBillingAt time.Time `json:"next_billing_at"`
}

func (u *UpdateSubscriptionBillingDateRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateSubscriptionBillingDateRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["next_billing_at"] = u.NextBillingAt.Format(time.RFC3339)
	return structMap
}

func (u *UpdateSubscriptionBillingDateRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		NextBillingAt string `json:"next_billing_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	NextBillingAtVal, err := time.Parse(time.RFC3339, temp.NextBillingAt)
	if err != nil {
		log.Fatalf("Cannot Parse next_billing_at as % s format.", time.RFC3339)
	}
	u.NextBillingAt = NextBillingAtVal
	return nil
}
