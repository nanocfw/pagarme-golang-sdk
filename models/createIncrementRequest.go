package models

import (
	"encoding/json"
)

// Request for creating a new increment
type CreateIncrementRequest struct {
	Value         float64 `json:"value"`
	IncrementType string  `json:"increment_type"`
	ItemId        string  `json:"item_id"`
	Cycles        *int    `json:"cycles,omitempty"`
	Description   *string `json:"description,omitempty"`
}

func (c *CreateIncrementRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateIncrementRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["value"] = c.Value
	structMap["increment_type"] = c.IncrementType
	structMap["item_id"] = c.ItemId
	if c.Cycles != nil {
		structMap["cycles"] = c.Cycles
	}
	if c.Description != nil {
		structMap["description"] = c.Description
	}
	return structMap
}

func (c *CreateIncrementRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Value         float64 `json:"value"`
		IncrementType string  `json:"increment_type"`
		ItemId        string  `json:"item_id"`
		Cycles        *int    `json:"cycles,omitempty"`
		Description   *string `json:"description,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Value = temp.Value
	c.IncrementType = temp.IncrementType
	c.ItemId = temp.ItemId
	c.Cycles = temp.Cycles
	c.Description = temp.Description
	return nil
}
