package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting an invoice item
type GetInvoiceItemResponse struct {
	Amount             types.Optional[int]                      `json:"amount"`
	Description        types.Optional[string]                   `json:"description"`
	PricingScheme      types.Optional[GetPricingSchemeResponse] `json:"pricing_scheme"`
	PriceBracket       types.Optional[GetPriceBracketResponse]  `json:"price_bracket"`
	Quantity           types.Optional[int]                      `json:"quantity"`
	Name               types.Optional[string]                   `json:"name"`
	SubscriptionItemId types.Optional[string]                   `json:"subscription_item_id"`
}

func (g *GetInvoiceItemResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetInvoiceItemResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	if g.PricingScheme.IsValueSet() {
		structMap["pricing_scheme"] = g.PricingScheme.Value()
	}
	if g.PriceBracket.IsValueSet() {
		structMap["price_bracket"] = g.PriceBracket.Value()
	}
	if g.Quantity.IsValueSet() {
		structMap["quantity"] = g.Quantity.Value()
	}
	if g.Name.IsValueSet() {
		structMap["name"] = g.Name.Value()
	}
	if g.SubscriptionItemId.IsValueSet() {
		structMap["subscription_item_id"] = g.SubscriptionItemId.Value()
	}
	return structMap
}

func (g *GetInvoiceItemResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount             types.Optional[int]                      `json:"amount"`
		Description        types.Optional[string]                   `json:"description"`
		PricingScheme      types.Optional[GetPricingSchemeResponse] `json:"pricing_scheme"`
		PriceBracket       types.Optional[GetPriceBracketResponse]  `json:"price_bracket"`
		Quantity           types.Optional[int]                      `json:"quantity"`
		Name               types.Optional[string]                   `json:"name"`
		SubscriptionItemId types.Optional[string]                   `json:"subscription_item_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Amount = temp.Amount
	g.Description = temp.Description
	g.PricingScheme = temp.PricingScheme
	g.PriceBracket = temp.PriceBracket
	g.Quantity = temp.Quantity
	g.Name = temp.Name
	g.SubscriptionItemId = temp.SubscriptionItemId
	return nil
}
