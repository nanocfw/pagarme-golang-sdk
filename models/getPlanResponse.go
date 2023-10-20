package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a plan
type GetPlanResponse struct {
	Id                  types.Optional[string]                `json:"id"`
	Name                types.Optional[string]                `json:"name"`
	Description         types.Optional[string]                `json:"description"`
	Url                 types.Optional[string]                `json:"url"`
	StatementDescriptor types.Optional[string]                `json:"statement_descriptor"`
	Interval            types.Optional[string]                `json:"interval"`
	IntervalCount       types.Optional[int]                   `json:"interval_count"`
	BillingType         types.Optional[string]                `json:"billing_type"`
	PaymentMethods      types.Optional[[]string]              `json:"payment_methods"`
	Installments        types.Optional[[]int]                 `json:"installments"`
	Status              types.Optional[string]                `json:"status"`
	Currency            types.Optional[string]                `json:"currency"`
	CreatedAt           types.Optional[time.Time]             `json:"created_at"`
	UpdatedAt           types.Optional[time.Time]             `json:"updated_at"`
	Items               types.Optional[[]GetPlanItemResponse] `json:"items"`
	BillingDays         types.Optional[[]int]                 `json:"billing_days"`
	Shippable           types.Optional[bool]                  `json:"shippable"`
	Metadata            types.Optional[map[string]string]     `json:"metadata"`
	TrialPeriodDays     types.Optional[int]                   `json:"trial_period_days"`
	MinimumPrice        types.Optional[int]                   `json:"minimum_price"`
	DeletedAt           types.Optional[time.Time]             `json:"deleted_at"`
}

func (g *GetPlanResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetPlanResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Name.IsValueSet() {
		structMap["name"] = g.Name.Value()
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	if g.Url.IsValueSet() {
		structMap["url"] = g.Url.Value()
	}
	if g.StatementDescriptor.IsValueSet() {
		structMap["statement_descriptor"] = g.StatementDescriptor.Value()
	}
	if g.Interval.IsValueSet() {
		structMap["interval"] = g.Interval.Value()
	}
	if g.IntervalCount.IsValueSet() {
		structMap["interval_count"] = g.IntervalCount.Value()
	}
	if g.BillingType.IsValueSet() {
		structMap["billing_type"] = g.BillingType.Value()
	}
	if g.PaymentMethods.IsValueSet() {
		structMap["payment_methods"] = g.PaymentMethods.Value()
	}
	if g.Installments.IsValueSet() {
		structMap["installments"] = g.Installments.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.Currency.IsValueSet() {
		structMap["currency"] = g.Currency.Value()
	}
	if g.CreatedAt.IsValueSet() {
		var CreatedAtVal *string = nil
		if g.CreatedAt.Value() != nil {
			val := g.CreatedAt.Value().Format(time.RFC3339)
			CreatedAtVal = &val
		}
		structMap["created_at"] = CreatedAtVal
	}
	if g.UpdatedAt.IsValueSet() {
		var UpdatedAtVal *string = nil
		if g.UpdatedAt.Value() != nil {
			val := g.UpdatedAt.Value().Format(time.RFC3339)
			UpdatedAtVal = &val
		}
		structMap["updated_at"] = UpdatedAtVal
	}
	if g.Items.IsValueSet() {
		structMap["items"] = g.Items.Value()
	}
	if g.BillingDays.IsValueSet() {
		structMap["billing_days"] = g.BillingDays.Value()
	}
	if g.Shippable.IsValueSet() {
		structMap["shippable"] = g.Shippable.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.TrialPeriodDays.IsValueSet() {
		structMap["trial_period_days"] = g.TrialPeriodDays.Value()
	}
	if g.MinimumPrice.IsValueSet() {
		structMap["minimum_price"] = g.MinimumPrice.Value()
	}
	if g.DeletedAt.IsValueSet() {
		var DeletedAtVal *string = nil
		if g.DeletedAt.Value() != nil {
			val := g.DeletedAt.Value().Format(time.RFC3339)
			DeletedAtVal = &val
		}
		structMap["deleted_at"] = DeletedAtVal
	}
	return structMap
}

func (g *GetPlanResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                  types.Optional[string]                `json:"id"`
		Name                types.Optional[string]                `json:"name"`
		Description         types.Optional[string]                `json:"description"`
		Url                 types.Optional[string]                `json:"url"`
		StatementDescriptor types.Optional[string]                `json:"statement_descriptor"`
		Interval            types.Optional[string]                `json:"interval"`
		IntervalCount       types.Optional[int]                   `json:"interval_count"`
		BillingType         types.Optional[string]                `json:"billing_type"`
		PaymentMethods      types.Optional[[]string]              `json:"payment_methods"`
		Installments        types.Optional[[]int]                 `json:"installments"`
		Status              types.Optional[string]                `json:"status"`
		Currency            types.Optional[string]                `json:"currency"`
		CreatedAt           types.Optional[string]                `json:"created_at"`
		UpdatedAt           types.Optional[string]                `json:"updated_at"`
		Items               types.Optional[[]GetPlanItemResponse] `json:"items"`
		BillingDays         types.Optional[[]int]                 `json:"billing_days"`
		Shippable           types.Optional[bool]                  `json:"shippable"`
		Metadata            types.Optional[map[string]string]     `json:"metadata"`
		TrialPeriodDays     types.Optional[int]                   `json:"trial_period_days"`
		MinimumPrice        types.Optional[int]                   `json:"minimum_price"`
		DeletedAt           types.Optional[string]                `json:"deleted_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Name = temp.Name
	g.Description = temp.Description
	g.Url = temp.Url
	g.StatementDescriptor = temp.StatementDescriptor
	g.Interval = temp.Interval
	g.IntervalCount = temp.IntervalCount
	g.BillingType = temp.BillingType
	g.PaymentMethods = temp.PaymentMethods
	g.Installments = temp.Installments
	g.Status = temp.Status
	g.Currency = temp.Currency
	g.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
	if temp.CreatedAt.Value() != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		g.CreatedAt.SetValue(&CreatedAtVal)
	}
	g.UpdatedAt.ShouldSetValue(temp.UpdatedAt.IsValueSet())
	if temp.UpdatedAt.Value() != nil {
		UpdatedAtVal, err := time.Parse(time.RFC3339, (*temp.UpdatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
		}
		g.UpdatedAt.SetValue(&UpdatedAtVal)
	}
	g.Items = temp.Items
	g.BillingDays = temp.BillingDays
	g.Shippable = temp.Shippable
	g.Metadata = temp.Metadata
	g.TrialPeriodDays = temp.TrialPeriodDays
	g.MinimumPrice = temp.MinimumPrice
	g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
	if temp.DeletedAt.Value() != nil {
		DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
		}
		g.DeletedAt.SetValue(&DeletedAtVal)
	}
	return nil
}
