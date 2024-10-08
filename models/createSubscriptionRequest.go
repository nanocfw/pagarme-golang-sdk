package models

import (
	"encoding/json"
	"log"
	"time"
)

// Request for creating a subcription
type CreateSubscriptionRequest struct {
	Customer             CreateCustomerRequest            `json:"customer"`
	Card                 CreateCardRequest                `json:"card"`
	Code                 string                           `json:"code"`
	PaymentMethod        string                           `json:"payment_method"`
	BillingType          string                           `json:"billing_type"`
	StatementDescriptor  string                           `json:"statement_descriptor"`
	Description          string                           `json:"description"`
	Currency             string                           `json:"currency"`
	Interval             string                           `json:"interval"`
	IntervalCount        int                              `json:"interval_count"`
	PricingScheme        CreatePricingSchemeRequest       `json:"pricing_scheme"`
	Items                []CreateSubscriptionItemRequest  `json:"items"`
	Shipping             CreateShippingRequest            `json:"shipping"`
	Discounts            []CreateDiscountRequest          `json:"discounts"`
	Metadata             map[string]string                `json:"metadata"`
	Setup                *CreateSetupRequest              `json:"setup,omitempty"`
	PlanId               *string                          `json:"plan_id,omitempty"`
	CustomerId           *string                          `json:"customer_id,omitempty"`
	CardId               *string                          `json:"card_id,omitempty"`
	BillingDay           *int                             `json:"billing_day,omitempty"`
	Installments         *int                             `json:"installments,omitempty"`
	StartAt              *time.Time                       `json:"start_at,omitempty"`
	MinimumPrice         *int                             `json:"minimum_price,omitempty"`
	Cycles               *int                             `json:"cycles,omitempty"`
	CardToken            *string                          `json:"card_token,omitempty"`
	GatewayAffiliationId *string                          `json:"gateway_affiliation_id,omitempty"`
	Quantity             *int                             `json:"quantity,omitempty"`
	BoletoDueDays        *int                             `json:"boleto_due_days,omitempty"`
	Increments           []CreateIncrementRequest         `json:"increments"`
	Period               *CreatePeriodRequest             `json:"period,omitempty"`
	Submerchant          *CreateSubMerchantRequest        `json:"submerchant,omitempty"`
	Split                *CreateSubscriptionSplitRequest  `json:"split,omitempty"`
	Boleto               *CreateSubscriptionBoletoRequest `json:"boleto,omitempty"`
}

func (c *CreateSubscriptionRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateSubscriptionRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["customer"] = c.Customer
	structMap["card"] = c.Card
	structMap["code"] = c.Code
	structMap["payment_method"] = c.PaymentMethod
	structMap["billing_type"] = c.BillingType
	structMap["statement_descriptor"] = c.StatementDescriptor
	structMap["description"] = c.Description
	structMap["currency"] = c.Currency
	structMap["interval"] = c.Interval
	structMap["interval_count"] = c.IntervalCount
	structMap["pricing_scheme"] = c.PricingScheme
	structMap["items"] = c.Items
	structMap["shipping"] = c.Shipping
	structMap["discounts"] = c.Discounts
	structMap["metadata"] = c.Metadata
	if c.Setup != nil {
		structMap["setup"] = c.Setup
	}
	if c.PlanId != nil {
		structMap["plan_id"] = c.PlanId
	}
	if c.CustomerId != nil {
		structMap["customer_id"] = c.CustomerId
	}
	if c.CardId != nil {
		structMap["card_id"] = c.CardId
	}
	if c.BillingDay != nil {
		structMap["billing_day"] = c.BillingDay
	}
	if c.Installments != nil {
		structMap["installments"] = c.Installments
	}
	if c.StartAt != nil {
		structMap["start_at"] = c.StartAt.Format(time.RFC3339)
	}
	if c.MinimumPrice != nil {
		structMap["minimum_price"] = c.MinimumPrice
	}
	if c.Cycles != nil {
		structMap["cycles"] = c.Cycles
	}
	if c.CardToken != nil {
		structMap["card_token"] = c.CardToken
	}
	if c.GatewayAffiliationId != nil {
		structMap["gateway_affiliation_id"] = c.GatewayAffiliationId
	}
	if c.Quantity != nil {
		structMap["quantity"] = c.Quantity
	}
	if c.BoletoDueDays != nil {
		structMap["boleto_due_days"] = c.BoletoDueDays
	}
	structMap["increments"] = c.Increments
	if c.Period != nil {
		structMap["period"] = c.Period
	}
	if c.Submerchant != nil {
		structMap["submerchant"] = c.Submerchant
	}
	if c.Split != nil {
		structMap["split"] = c.Split
	}
	if c.Boleto != nil {
		structMap["boleto"] = c.Boleto
	}
	return structMap
}

func (c *CreateSubscriptionRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Customer             CreateCustomerRequest            `json:"customer"`
		Card                 CreateCardRequest                `json:"card"`
		Code                 string                           `json:"code"`
		PaymentMethod        string                           `json:"payment_method"`
		BillingType          string                           `json:"billing_type"`
		StatementDescriptor  string                           `json:"statement_descriptor"`
		Description          string                           `json:"description"`
		Currency             string                           `json:"currency"`
		Interval             string                           `json:"interval"`
		IntervalCount        int                              `json:"interval_count"`
		PricingScheme        CreatePricingSchemeRequest       `json:"pricing_scheme"`
		Items                []CreateSubscriptionItemRequest  `json:"items"`
		Shipping             CreateShippingRequest            `json:"shipping"`
		Discounts            []CreateDiscountRequest          `json:"discounts"`
		Metadata             map[string]string                `json:"metadata"`
		Setup                *CreateSetupRequest              `json:"setup,omitempty"`
		PlanId               *string                          `json:"plan_id,omitempty"`
		CustomerId           *string                          `json:"customer_id,omitempty"`
		CardId               *string                          `json:"card_id,omitempty"`
		BillingDay           *int                             `json:"billing_day,omitempty"`
		Installments         *int                             `json:"installments,omitempty"`
		StartAt              *string                          `json:"start_at,omitempty"`
		MinimumPrice         *int                             `json:"minimum_price,omitempty"`
		Cycles               *int                             `json:"cycles,omitempty"`
		CardToken            *string                          `json:"card_token,omitempty"`
		GatewayAffiliationId *string                          `json:"gateway_affiliation_id,omitempty"`
		Quantity             *int                             `json:"quantity,omitempty"`
		BoletoDueDays        *int                             `json:"boleto_due_days,omitempty"`
		Increments           []CreateIncrementRequest         `json:"increments"`
		Period               *CreatePeriodRequest             `json:"period,omitempty"`
		Submerchant          *CreateSubMerchantRequest        `json:"submerchant,omitempty"`
		Split                *CreateSubscriptionSplitRequest  `json:"split,omitempty"`
		Boleto               *CreateSubscriptionBoletoRequest `json:"boleto,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Customer = temp.Customer
	c.Card = temp.Card
	c.Code = temp.Code
	c.PaymentMethod = temp.PaymentMethod
	c.BillingType = temp.BillingType
	c.StatementDescriptor = temp.StatementDescriptor
	c.Description = temp.Description
	c.Currency = temp.Currency
	c.Interval = temp.Interval
	c.IntervalCount = temp.IntervalCount
	c.PricingScheme = temp.PricingScheme
	c.Items = temp.Items
	c.Shipping = temp.Shipping
	c.Discounts = temp.Discounts
	c.Metadata = temp.Metadata
	c.Setup = temp.Setup
	c.PlanId = temp.PlanId
	c.CustomerId = temp.CustomerId
	c.CardId = temp.CardId
	c.BillingDay = temp.BillingDay
	c.Installments = temp.Installments
	if temp.StartAt != nil {
		StartAtVal, err := time.Parse(time.RFC3339, *temp.StartAt)
		if err != nil {
			log.Fatalf("Cannot Parse start_at as % s format.", time.RFC3339)
		}
		c.StartAt = &StartAtVal
	}
	c.MinimumPrice = temp.MinimumPrice
	c.Cycles = temp.Cycles
	c.CardToken = temp.CardToken
	c.GatewayAffiliationId = temp.GatewayAffiliationId
	c.Quantity = temp.Quantity
	c.BoletoDueDays = temp.BoletoDueDays
	c.Increments = temp.Increments
	c.Period = temp.Period
	c.Submerchant = temp.Submerchant
	c.Split = temp.Split
	c.Boleto = temp.Boleto
	return nil
}
