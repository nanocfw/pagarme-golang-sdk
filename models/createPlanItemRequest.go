package models

import (
	"encoding/json"
)

// Request for creating a plan item
type CreatePlanItemRequest struct {
	Name          string                     `json:"name"`
	PricingScheme CreatePricingSchemeRequest `json:"pricing_scheme"`
	Id            string                     `json:"id"`
	Description   string                     `json:"description"`
	Cycles        *int                       `json:"cycles,omitempty"`
	Quantity      *int                       `json:"quantity,omitempty"`
}

func (c *CreatePlanItemRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreatePlanItemRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["name"] = c.Name
	structMap["pricing_scheme"] = c.PricingScheme
	structMap["id"] = c.Id
	structMap["description"] = c.Description
	if c.Cycles != nil {
		structMap["cycles"] = c.Cycles
	}
	if c.Quantity != nil {
		structMap["quantity"] = c.Quantity
	}
	return structMap
}

func (c *CreatePlanItemRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name          string                     `json:"name"`
		PricingScheme CreatePricingSchemeRequest `json:"pricing_scheme"`
		Id            string                     `json:"id"`
		Description   string                     `json:"description"`
		Cycles        *int                       `json:"cycles,omitempty"`
		Quantity      *int                       `json:"quantity,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Name = temp.Name
	c.PricingScheme = temp.PricingScheme
	c.Id = temp.Id
	c.Description = temp.Description
	c.Cycles = temp.Cycles
	c.Quantity = temp.Quantity
	return nil
}
