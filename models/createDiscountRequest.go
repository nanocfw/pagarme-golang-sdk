package models

import (
	"encoding/json"
)

// Request for creating a new discount
type CreateDiscountRequest struct {
	Value        float64 `json:"value"`
	DiscountType string  `json:"discount_type"`
	ItemId       string  `json:"item_id"`
	Cycles       *int    `json:"cycles,omitempty"`
	Description  *string `json:"description,omitempty"`
}

func (c *CreateDiscountRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateDiscountRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["value"] = c.Value
	structMap["discount_type"] = c.DiscountType
	structMap["item_id"] = c.ItemId
	if c.Cycles != nil {
		structMap["cycles"] = c.Cycles
	}
	if c.Description != nil {
		structMap["description"] = c.Description
	}
	return structMap
}

func (c *CreateDiscountRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Value        float64 `json:"value"`
		DiscountType string  `json:"discount_type"`
		ItemId       string  `json:"item_id"`
		Cycles       *int    `json:"cycles,omitempty"`
		Description  *string `json:"description,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Value = temp.Value
	c.DiscountType = temp.DiscountType
	c.ItemId = temp.ItemId
	c.Cycles = temp.Cycles
	c.Description = temp.Description
	return nil
}
