package models

import (
	"encoding/json"
)

// Request for creating a plan
type CreatePlanRequest struct {
	Name                string                     `json:"name"`
	Description         string                     `json:"description"`
	StatementDescriptor string                     `json:"statement_descriptor"`
	Items               []CreatePlanItemRequest    `json:"items"`
	Shippable           bool                       `json:"shippable"`
	PaymentMethods      []string                   `json:"payment_methods"`
	Installments        []int                      `json:"installments"`
	Currency            string                     `json:"currency"`
	Interval            string                     `json:"interval"`
	IntervalCount       int                        `json:"interval_count"`
	BillingDays         []int                      `json:"billing_days"`
	BillingType         string                     `json:"billing_type"`
	PricingScheme       CreatePricingSchemeRequest `json:"pricing_scheme"`
	Metadata            map[string]string          `json:"metadata"`
	MinimumPrice        *int                       `json:"minimum_price,omitempty"`
	Cycles              *int                       `json:"cycles,omitempty"`
	Quantity            *int                       `json:"quantity,omitempty"`
	TrialPeriodDays     *int                       `json:"trial_period_days,omitempty"`
}

func (c *CreatePlanRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreatePlanRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["name"] = c.Name
	structMap["description"] = c.Description
	structMap["statement_descriptor"] = c.StatementDescriptor
	structMap["items"] = c.Items
	structMap["shippable"] = c.Shippable
	structMap["payment_methods"] = c.PaymentMethods
	structMap["installments"] = c.Installments
	structMap["currency"] = c.Currency
	structMap["interval"] = c.Interval
	structMap["interval_count"] = c.IntervalCount
	structMap["billing_days"] = c.BillingDays
	structMap["billing_type"] = c.BillingType
	structMap["pricing_scheme"] = c.PricingScheme
	structMap["metadata"] = c.Metadata
	if c.MinimumPrice != nil {
		structMap["minimum_price"] = c.MinimumPrice
	}
	if c.Cycles != nil {
		structMap["cycles"] = c.Cycles
	}
	if c.Quantity != nil {
		structMap["quantity"] = c.Quantity
	}
	if c.TrialPeriodDays != nil {
		structMap["trial_period_days"] = c.TrialPeriodDays
	}
	return structMap
}

func (c *CreatePlanRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name                string                     `json:"name"`
		Description         string                     `json:"description"`
		StatementDescriptor string                     `json:"statement_descriptor"`
		Items               []CreatePlanItemRequest    `json:"items"`
		Shippable           bool                       `json:"shippable"`
		PaymentMethods      []string                   `json:"payment_methods"`
		Installments        []int                      `json:"installments"`
		Currency            string                     `json:"currency"`
		Interval            string                     `json:"interval"`
		IntervalCount       int                        `json:"interval_count"`
		BillingDays         []int                      `json:"billing_days"`
		BillingType         string                     `json:"billing_type"`
		PricingScheme       CreatePricingSchemeRequest `json:"pricing_scheme"`
		Metadata            map[string]string          `json:"metadata"`
		MinimumPrice        *int                       `json:"minimum_price,omitempty"`
		Cycles              *int                       `json:"cycles,omitempty"`
		Quantity            *int                       `json:"quantity,omitempty"`
		TrialPeriodDays     *int                       `json:"trial_period_days,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Name = temp.Name
	c.Description = temp.Description
	c.StatementDescriptor = temp.StatementDescriptor
	c.Items = temp.Items
	c.Shippable = temp.Shippable
	c.PaymentMethods = temp.PaymentMethods
	c.Installments = temp.Installments
	c.Currency = temp.Currency
	c.Interval = temp.Interval
	c.IntervalCount = temp.IntervalCount
	c.BillingDays = temp.BillingDays
	c.BillingType = temp.BillingType
	c.PricingScheme = temp.PricingScheme
	c.Metadata = temp.Metadata
	c.MinimumPrice = temp.MinimumPrice
	c.Cycles = temp.Cycles
	c.Quantity = temp.Quantity
	c.TrialPeriodDays = temp.TrialPeriodDays
	return nil
}
