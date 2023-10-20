package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Checkout Payment Settings Response
type GetCheckoutPaymentSettingsResponse struct {
	SuccessUrl             types.Optional[string]              `json:"success_url"`
	PaymentUrl             types.Optional[string]              `json:"payment_url"`
	AcceptedPaymentMethods types.Optional[[]string]            `json:"accepted_payment_methods"`
	Status                 types.Optional[string]              `json:"status"`
	Customer               types.Optional[GetCustomerResponse] `json:"customer"`
	Amount                 types.Optional[int]                 `json:"amount"`
	DefaultPaymentMethod   types.Optional[string]              `json:"default_payment_method"`
	GatewayAffiliationId   types.Optional[string]              `json:"gateway_affiliation_id"`
}

func (g *GetCheckoutPaymentSettingsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetCheckoutPaymentSettingsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.SuccessUrl.IsValueSet() {
		structMap["success_url"] = g.SuccessUrl.Value()
	}
	if g.PaymentUrl.IsValueSet() {
		structMap["payment_url"] = g.PaymentUrl.Value()
	}
	if g.AcceptedPaymentMethods.IsValueSet() {
		structMap["accepted_payment_methods"] = g.AcceptedPaymentMethods.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.Customer.IsValueSet() {
		structMap["customer"] = g.Customer.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.DefaultPaymentMethod.IsValueSet() {
		structMap["default_payment_method"] = g.DefaultPaymentMethod.Value()
	}
	if g.GatewayAffiliationId.IsValueSet() {
		structMap["gateway_affiliation_id"] = g.GatewayAffiliationId.Value()
	}
	return structMap
}

func (g *GetCheckoutPaymentSettingsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SuccessUrl             types.Optional[string]              `json:"success_url"`
		PaymentUrl             types.Optional[string]              `json:"payment_url"`
		AcceptedPaymentMethods types.Optional[[]string]            `json:"accepted_payment_methods"`
		Status                 types.Optional[string]              `json:"status"`
		Customer               types.Optional[GetCustomerResponse] `json:"customer"`
		Amount                 types.Optional[int]                 `json:"amount"`
		DefaultPaymentMethod   types.Optional[string]              `json:"default_payment_method"`
		GatewayAffiliationId   types.Optional[string]              `json:"gateway_affiliation_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.SuccessUrl = temp.SuccessUrl
	g.PaymentUrl = temp.PaymentUrl
	g.AcceptedPaymentMethods = temp.AcceptedPaymentMethods
	g.Status = temp.Status
	g.Customer = temp.Customer
	g.Amount = temp.Amount
	g.DefaultPaymentMethod = temp.DefaultPaymentMethod
	g.GatewayAffiliationId = temp.GatewayAffiliationId
	return nil
}
