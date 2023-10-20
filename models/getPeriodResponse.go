package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a period
type GetPeriodResponse struct {
	StartAt      types.Optional[time.Time]               `json:"start_at"`
	EndAt        types.Optional[time.Time]               `json:"end_at"`
	Id           types.Optional[string]                  `json:"id"`
	BillingAt    types.Optional[time.Time]               `json:"billing_at"`
	Subscription types.Optional[GetSubscriptionResponse] `json:"subscription"`
	Status       types.Optional[string]                  `json:"status"`
	Duration     types.Optional[int]                     `json:"duration"`
	CreatedAt    types.Optional[string]                  `json:"created_at"`
	UpdatedAt    types.Optional[string]                  `json:"updated_at"`
	Cycle        types.Optional[int]                     `json:"cycle"`
}

func (g *GetPeriodResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetPeriodResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.StartAt.IsValueSet() {
		var StartAtVal *string = nil
		if g.StartAt.Value() != nil {
			val := g.StartAt.Value().Format(time.RFC3339)
			StartAtVal = &val
		}
		structMap["start_at"] = StartAtVal
	}
	if g.EndAt.IsValueSet() {
		var EndAtVal *string = nil
		if g.EndAt.Value() != nil {
			val := g.EndAt.Value().Format(time.RFC3339)
			EndAtVal = &val
		}
		structMap["end_at"] = EndAtVal
	}
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.BillingAt.IsValueSet() {
		var BillingAtVal *string = nil
		if g.BillingAt.Value() != nil {
			val := g.BillingAt.Value().Format(time.RFC3339)
			BillingAtVal = &val
		}
		structMap["billing_at"] = BillingAtVal
	}
	if g.Subscription.IsValueSet() {
		structMap["subscription"] = g.Subscription.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.Duration.IsValueSet() {
		structMap["duration"] = g.Duration.Value()
	}
	if g.CreatedAt.IsValueSet() {
		structMap["created_at"] = g.CreatedAt.Value()
	}
	if g.UpdatedAt.IsValueSet() {
		structMap["updated_at"] = g.UpdatedAt.Value()
	}
	if g.Cycle.IsValueSet() {
		structMap["cycle"] = g.Cycle.Value()
	}
	return structMap
}

func (g *GetPeriodResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StartAt      types.Optional[string]                  `json:"start_at"`
		EndAt        types.Optional[string]                  `json:"end_at"`
		Id           types.Optional[string]                  `json:"id"`
		BillingAt    types.Optional[string]                  `json:"billing_at"`
		Subscription types.Optional[GetSubscriptionResponse] `json:"subscription"`
		Status       types.Optional[string]                  `json:"status"`
		Duration     types.Optional[int]                     `json:"duration"`
		CreatedAt    types.Optional[string]                  `json:"created_at"`
		UpdatedAt    types.Optional[string]                  `json:"updated_at"`
		Cycle        types.Optional[int]                     `json:"cycle"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.StartAt.ShouldSetValue(temp.StartAt.IsValueSet())
	if temp.StartAt.Value() != nil {
		StartAtVal, err := time.Parse(time.RFC3339, (*temp.StartAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse start_at as % s format.", time.RFC3339)
		}
		g.StartAt.SetValue(&StartAtVal)
	}
	g.EndAt.ShouldSetValue(temp.EndAt.IsValueSet())
	if temp.EndAt.Value() != nil {
		EndAtVal, err := time.Parse(time.RFC3339, (*temp.EndAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse end_at as % s format.", time.RFC3339)
		}
		g.EndAt.SetValue(&EndAtVal)
	}
	g.Id = temp.Id
	g.BillingAt.ShouldSetValue(temp.BillingAt.IsValueSet())
	if temp.BillingAt.Value() != nil {
		BillingAtVal, err := time.Parse(time.RFC3339, (*temp.BillingAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse billing_at as % s format.", time.RFC3339)
		}
		g.BillingAt.SetValue(&BillingAtVal)
	}
	g.Subscription = temp.Subscription
	g.Status = temp.Status
	g.Duration = temp.Duration
	g.CreatedAt = temp.CreatedAt
	g.UpdatedAt = temp.UpdatedAt
	g.Cycle = temp.Cycle
	return nil
}
