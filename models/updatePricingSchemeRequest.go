package models

import (
	"encoding/json"
)

// Request for updating a pricing scheme
type UpdatePricingSchemeRequest struct {
	SchemeType    string                      `json:"scheme_type"`
	PriceBrackets []UpdatePriceBracketRequest `json:"price_brackets"`
	Price         *int                        `json:"price,omitempty"`
	MinimumPrice  *int                        `json:"minimum_price,omitempty"`
	Percentage    *float64                    `json:"percentage,omitempty"`
}

func (u *UpdatePricingSchemeRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdatePricingSchemeRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["scheme_type"] = u.SchemeType
	structMap["price_brackets"] = u.PriceBrackets
	if u.Price != nil {
		structMap["price"] = u.Price
	}
	if u.MinimumPrice != nil {
		structMap["minimum_price"] = u.MinimumPrice
	}
	if u.Percentage != nil {
		structMap["percentage"] = u.Percentage
	}
	return structMap
}

func (u *UpdatePricingSchemeRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SchemeType    string                      `json:"scheme_type"`
		PriceBrackets []UpdatePriceBracketRequest `json:"price_brackets"`
		Price         *int                        `json:"price,omitempty"`
		MinimumPrice  *int                        `json:"minimum_price,omitempty"`
		Percentage    *float64                    `json:"percentage,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.SchemeType = temp.SchemeType
	u.PriceBrackets = temp.PriceBrackets
	u.Price = temp.Price
	u.MinimumPrice = temp.MinimumPrice
	u.Percentage = temp.Percentage
	return nil
}
