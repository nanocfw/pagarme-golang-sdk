package models

import (
	"encoding/json"
)

// Request for creating a price bracket
type CreatePriceBracketRequest struct {
	StartQuantity int  `json:"start_quantity"`
	Price         int  `json:"price"`
	EndQuantity   *int `json:"end_quantity,omitempty"`
	OveragePrice  *int `json:"overage_price,omitempty"`
}

func (c *CreatePriceBracketRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreatePriceBracketRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["start_quantity"] = c.StartQuantity
	structMap["price"] = c.Price
	if c.EndQuantity != nil {
		structMap["end_quantity"] = c.EndQuantity
	}
	if c.OveragePrice != nil {
		structMap["overage_price"] = c.OveragePrice
	}
	return structMap
}

func (c *CreatePriceBracketRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StartQuantity int  `json:"start_quantity"`
		Price         int  `json:"price"`
		EndQuantity   *int `json:"end_quantity,omitempty"`
		OveragePrice  *int `json:"overage_price,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.StartQuantity = temp.StartQuantity
	c.Price = temp.Price
	c.EndQuantity = temp.EndQuantity
	c.OveragePrice = temp.OveragePrice
	return nil
}
