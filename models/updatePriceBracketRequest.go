package models

import (
	"encoding/json"
)

// Request for updating a price bracket
type UpdatePriceBracketRequest struct {
	StartQuantity int  `json:"start_quantity"`
	Price         int  `json:"price"`
	EndQuantity   *int `json:"end_quantity,omitempty"`
	OveragePrice  *int `json:"overage_price,omitempty"`
}

func (u *UpdatePriceBracketRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdatePriceBracketRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["start_quantity"] = u.StartQuantity
	structMap["price"] = u.Price
	if u.EndQuantity != nil {
		structMap["end_quantity"] = u.EndQuantity
	}
	if u.OveragePrice != nil {
		structMap["overage_price"] = u.OveragePrice
	}
	return structMap
}

func (u *UpdatePriceBracketRequest) UnmarshalJSON(input []byte) error {
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

	u.StartQuantity = temp.StartQuantity
	u.Price = temp.Price
	u.EndQuantity = temp.EndQuantity
	u.OveragePrice = temp.OveragePrice
	return nil
}
