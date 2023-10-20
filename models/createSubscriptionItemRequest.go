package models

import (
	"encoding/json"
)

// Request for creating a new subscription item
type CreateSubscriptionItemRequest struct {
	Description   string                     `json:"description"`
	PricingScheme CreatePricingSchemeRequest `json:"pricing_scheme"`
	Id            string                     `json:"id"`
	PlanItemId    string                     `json:"plan_item_id"`
	Discounts     []CreateDiscountRequest    `json:"discounts"`
	Name          string                     `json:"name"`
	Cycles        *int                       `json:"cycles,omitempty"`
	Quantity      *int                       `json:"quantity,omitempty"`
	MinimumPrice  *int                       `json:"minimum_price,omitempty"`
}

func (c *CreateSubscriptionItemRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateSubscriptionItemRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["description"] = c.Description
	structMap["pricing_scheme"] = c.PricingScheme
	structMap["id"] = c.Id
	structMap["plan_item_id"] = c.PlanItemId
	structMap["discounts"] = c.Discounts
	structMap["name"] = c.Name
	if c.Cycles != nil {
		structMap["cycles"] = c.Cycles
	}
	if c.Quantity != nil {
		structMap["quantity"] = c.Quantity
	}
	if c.MinimumPrice != nil {
		structMap["minimum_price"] = c.MinimumPrice
	}
	return structMap
}

func (c *CreateSubscriptionItemRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Description   string                     `json:"description"`
		PricingScheme CreatePricingSchemeRequest `json:"pricing_scheme"`
		Id            string                     `json:"id"`
		PlanItemId    string                     `json:"plan_item_id"`
		Discounts     []CreateDiscountRequest    `json:"discounts"`
		Name          string                     `json:"name"`
		Cycles        *int                       `json:"cycles,omitempty"`
		Quantity      *int                       `json:"quantity,omitempty"`
		MinimumPrice  *int                       `json:"minimum_price,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Description = temp.Description
	c.PricingScheme = temp.PricingScheme
	c.Id = temp.Id
	c.PlanItemId = temp.PlanItemId
	c.Discounts = temp.Discounts
	c.Name = temp.Name
	c.Cycles = temp.Cycles
	c.Quantity = temp.Quantity
	c.MinimumPrice = temp.MinimumPrice
	return nil
}
