package models

import (
	"encoding/json"
)

// Request for creating a pricing scheme
type CreatePricingSchemeRequest struct {
	SchemeType    string                       `json:"scheme_type"`
	PriceBrackets *[]CreatePriceBracketRequest `json:"price_brackets,omitempty"`
	Price         *int                         `json:"price,omitempty"`
	MinimumPrice  *int                         `json:"minimum_price,omitempty"`
	Percentage    *float64                     `json:"percentage,omitempty"`
}

func (c *CreatePricingSchemeRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreatePricingSchemeRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["scheme_type"] = c.SchemeType
	if c.PriceBrackets != nil {
		structMap["price_brackets"] = c.PriceBrackets
	}
	if c.Price != nil {
		structMap["price"] = c.Price
	}
	if c.MinimumPrice != nil {
		structMap["minimum_price"] = c.MinimumPrice
	}
	if c.Percentage != nil {
		structMap["percentage"] = c.Percentage
	}
	return structMap
}

func (c *CreatePricingSchemeRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SchemeType    string                       `json:"scheme_type"`
		PriceBrackets *[]CreatePriceBracketRequest `json:"price_brackets,omitempty"`
		Price         *int                         `json:"price,omitempty"`
		MinimumPrice  *int                         `json:"minimum_price,omitempty"`
		Percentage    *float64                     `json:"percentage,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.SchemeType = temp.SchemeType
	c.PriceBrackets = temp.PriceBrackets
	c.Price = temp.Price
	c.MinimumPrice = temp.MinimumPrice
	c.Percentage = temp.Percentage
	return nil
}
