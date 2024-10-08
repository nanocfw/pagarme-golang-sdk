package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

type GetSubscriptionResponse struct {
	Id                   types.Optional[string]                        `json:"id"`
	Code                 types.Optional[string]                        `json:"code"`
	StartAt              types.Optional[time.Time]                     `json:"start_at"`
	Interval             types.Optional[string]                        `json:"interval"`
	IntervalCount        types.Optional[int]                           `json:"interval_count"`
	BillingType          types.Optional[string]                        `json:"billing_type"`
	CurrentCycle         types.Optional[GetPeriodResponse]             `json:"current_cycle"`
	PaymentMethod        types.Optional[string]                        `json:"payment_method"`
	Currency             types.Optional[string]                        `json:"currency"`
	Installments         types.Optional[int]                           `json:"installments"`
	Status               types.Optional[string]                        `json:"status"`
	CreatedAt            types.Optional[time.Time]                     `json:"created_at"`
	UpdatedAt            types.Optional[time.Time]                     `json:"updated_at"`
	Customer             types.Optional[GetCustomerResponse]           `json:"customer"`
	Card                 types.Optional[GetCardResponse]               `json:"card"`
	Items                types.Optional[[]GetSubscriptionItemResponse] `json:"items"`
	StatementDescriptor  types.Optional[string]                        `json:"statement_descriptor"`
	Metadata             types.Optional[map[string]string]             `json:"metadata"`
	Setup                types.Optional[GetSetupResponse]              `json:"setup"`
	GatewayAffiliationId types.Optional[string]                        `json:"gateway_affiliation_id"`
	NextBillingAt        types.Optional[time.Time]                     `json:"next_billing_at"`
	BillingDay           types.Optional[int]                           `json:"billing_day"`
	MinimumPrice         types.Optional[int]                           `json:"minimum_price"`
	CanceledAt           types.Optional[time.Time]                     `json:"canceled_at"`
	Discounts            types.Optional[[]GetDiscountResponse]         `json:"discounts"`
	Increments           types.Optional[[]GetIncrementResponse]        `json:"increments"`
	BoletoDueDays        types.Optional[int]                           `json:"boleto_due_days"`
	Split                types.Optional[GetSubscriptionSplitResponse]  `json:"split"`
	Boleto               types.Optional[GetSubscriptionBoletoResponse] `json:"boleto"`
	ManualBilling        types.Optional[bool]                          `json:"manual_billing"`
}

func (g *GetSubscriptionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetSubscriptionResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Code.IsValueSet() {
		structMap["code"] = g.Code.Value()
	}
	if g.StartAt.IsValueSet() {
		var StartAtVal *string = nil
		if g.StartAt.Value() != nil {
			val := g.StartAt.Value().Format(time.RFC3339)
			StartAtVal = &val
		}
		structMap["start_at"] = StartAtVal
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
	if g.CurrentCycle.IsValueSet() {
		structMap["current_cycle"] = g.CurrentCycle.Value()
	}
	if g.PaymentMethod.IsValueSet() {
		structMap["payment_method"] = g.PaymentMethod.Value()
	}
	if g.Currency.IsValueSet() {
		structMap["currency"] = g.Currency.Value()
	}
	if g.Installments.IsValueSet() {
		structMap["installments"] = g.Installments.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
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
	if g.Customer.IsValueSet() {
		structMap["customer"] = g.Customer.Value()
	}
	if g.Card.IsValueSet() {
		structMap["card"] = g.Card.Value()
	}
	if g.Items.IsValueSet() {
		structMap["items"] = g.Items.Value()
	}
	if g.StatementDescriptor.IsValueSet() {
		structMap["statement_descriptor"] = g.StatementDescriptor.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.Setup.IsValueSet() {
		structMap["setup"] = g.Setup.Value()
	}
	if g.GatewayAffiliationId.IsValueSet() {
		structMap["gateway_affiliation_id"] = g.GatewayAffiliationId.Value()
	}
	if g.NextBillingAt.IsValueSet() {
		var NextBillingAtVal *string = nil
		if g.NextBillingAt.Value() != nil {
			val := g.NextBillingAt.Value().Format(time.RFC3339)
			NextBillingAtVal = &val
		}
		structMap["next_billing_at"] = NextBillingAtVal
	}
	if g.BillingDay.IsValueSet() {
		structMap["billing_day"] = g.BillingDay.Value()
	}
	if g.MinimumPrice.IsValueSet() {
		structMap["minimum_price"] = g.MinimumPrice.Value()
	}
	if g.CanceledAt.IsValueSet() {
		var CanceledAtVal *string = nil
		if g.CanceledAt.Value() != nil {
			val := g.CanceledAt.Value().Format(time.RFC3339)
			CanceledAtVal = &val
		}
		structMap["canceled_at"] = CanceledAtVal
	}
	if g.Discounts.IsValueSet() {
		structMap["discounts"] = g.Discounts.Value()
	}
	if g.Increments.IsValueSet() {
		structMap["increments"] = g.Increments.Value()
	}
	if g.BoletoDueDays.IsValueSet() {
		structMap["boleto_due_days"] = g.BoletoDueDays.Value()
	}
	if g.Split.IsValueSet() {
		structMap["split"] = g.Split.Value()
	}
	if g.Boleto.IsValueSet() {
		structMap["boleto"] = g.Boleto.Value()
	}
	if g.ManualBilling.IsValueSet() {
		structMap["manual_billing"] = g.ManualBilling.Value()
	}
	return structMap
}

func (g *GetSubscriptionResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                   types.Optional[string]                        `json:"id"`
		Code                 types.Optional[string]                        `json:"code"`
		StartAt              types.Optional[string]                        `json:"start_at"`
		Interval             types.Optional[string]                        `json:"interval"`
		IntervalCount        types.Optional[int]                           `json:"interval_count"`
		BillingType          types.Optional[string]                        `json:"billing_type"`
		CurrentCycle         types.Optional[GetPeriodResponse]             `json:"current_cycle"`
		PaymentMethod        types.Optional[string]                        `json:"payment_method"`
		Currency             types.Optional[string]                        `json:"currency"`
		Installments         types.Optional[int]                           `json:"installments"`
		Status               types.Optional[string]                        `json:"status"`
		CreatedAt            types.Optional[string]                        `json:"created_at"`
		UpdatedAt            types.Optional[string]                        `json:"updated_at"`
		Customer             types.Optional[GetCustomerResponse]           `json:"customer"`
		Card                 types.Optional[GetCardResponse]               `json:"card"`
		Items                types.Optional[[]GetSubscriptionItemResponse] `json:"items"`
		StatementDescriptor  types.Optional[string]                        `json:"statement_descriptor"`
		Metadata             types.Optional[map[string]string]             `json:"metadata"`
		Setup                types.Optional[GetSetupResponse]              `json:"setup"`
		GatewayAffiliationId types.Optional[string]                        `json:"gateway_affiliation_id"`
		NextBillingAt        types.Optional[string]                        `json:"next_billing_at"`
		BillingDay           types.Optional[int]                           `json:"billing_day"`
		MinimumPrice         types.Optional[int]                           `json:"minimum_price"`
		CanceledAt           types.Optional[string]                        `json:"canceled_at"`
		Discounts            types.Optional[[]GetDiscountResponse]         `json:"discounts"`
		Increments           types.Optional[[]GetIncrementResponse]        `json:"increments"`
		BoletoDueDays        types.Optional[int]                           `json:"boleto_due_days"`
		Split                types.Optional[GetSubscriptionSplitResponse]  `json:"split"`
		Boleto               types.Optional[GetSubscriptionBoletoResponse] `json:"boleto"`
		ManualBilling        types.Optional[bool]                          `json:"manual_billing"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Code = temp.Code
	g.StartAt.ShouldSetValue(temp.StartAt.IsValueSet())
	if temp.StartAt.Value() != nil {
		StartAtVal, err := time.Parse(time.RFC3339, (*temp.StartAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse start_at as % s format.", time.RFC3339)
		}
		g.StartAt.SetValue(&StartAtVal)
	}
	g.Interval = temp.Interval
	g.IntervalCount = temp.IntervalCount
	g.BillingType = temp.BillingType
	g.CurrentCycle = temp.CurrentCycle
	g.PaymentMethod = temp.PaymentMethod
	g.Currency = temp.Currency
	g.Installments = temp.Installments
	g.Status = temp.Status
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
	g.Customer = temp.Customer
	g.Card = temp.Card
	g.Items = temp.Items
	g.StatementDescriptor = temp.StatementDescriptor
	g.Metadata = temp.Metadata
	g.Setup = temp.Setup
	g.GatewayAffiliationId = temp.GatewayAffiliationId
	g.NextBillingAt.ShouldSetValue(temp.NextBillingAt.IsValueSet())
	if temp.NextBillingAt.Value() != nil {
		NextBillingAtVal, err := time.Parse(time.RFC3339, (*temp.NextBillingAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse next_billing_at as % s format.", time.RFC3339)
		}
		g.NextBillingAt.SetValue(&NextBillingAtVal)
	}
	g.BillingDay = temp.BillingDay
	g.MinimumPrice = temp.MinimumPrice
	g.CanceledAt.ShouldSetValue(temp.CanceledAt.IsValueSet())
	if temp.CanceledAt.Value() != nil {
		CanceledAtVal, err := time.Parse(time.RFC3339, (*temp.CanceledAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse canceled_at as % s format.", time.RFC3339)
		}
		g.CanceledAt.SetValue(&CanceledAtVal)
	}
	g.Discounts = temp.Discounts
	g.Increments = temp.Increments
	g.BoletoDueDays = temp.BoletoDueDays
	g.Split = temp.Split
	g.Boleto = temp.Boleto
	g.ManualBilling = temp.ManualBilling
	return nil
}
