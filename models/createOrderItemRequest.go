package models

import (
	"encoding/json"
)

// Request for creating an order item
type CreateOrderItemRequest struct {
	Amount      int     `json:"amount"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Category    string  `json:"category"`
	Code        *string `json:"code,omitempty"`
}

func (c *CreateOrderItemRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateOrderItemRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = c.Amount
	structMap["description"] = c.Description
	structMap["quantity"] = c.Quantity
	structMap["category"] = c.Category
	if c.Code != nil {
		structMap["code"] = c.Code
	}
	return structMap
}

func (c *CreateOrderItemRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount      int     `json:"amount"`
		Description string  `json:"description"`
		Quantity    int     `json:"quantity"`
		Category    string  `json:"category"`
		Code        *string `json:"code,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Amount = temp.Amount
	c.Description = temp.Description
	c.Quantity = temp.Quantity
	c.Category = temp.Category
	c.Code = temp.Code
	return nil
}
