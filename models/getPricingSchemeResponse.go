package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a pricing scheme
type GetPricingSchemeResponse struct {
	Price         types.Optional[int]                       `json:"price"`
	SchemeType    types.Optional[string]                    `json:"scheme_type"`
	PriceBrackets types.Optional[[]GetPriceBracketResponse] `json:"price_brackets"`
	MinimumPrice  types.Optional[int]                       `json:"minimum_price"`
	Percentage    types.Optional[float64]                   `json:"percentage"`
}

func (g *GetPricingSchemeResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetPricingSchemeResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Price.IsValueSet() {
		structMap["price"] = g.Price.Value()
	}
	if g.SchemeType.IsValueSet() {
		structMap["scheme_type"] = g.SchemeType.Value()
	}
	if g.PriceBrackets.IsValueSet() {
		structMap["price_brackets"] = g.PriceBrackets.Value()
	}
	if g.MinimumPrice.IsValueSet() {
		structMap["minimum_price"] = g.MinimumPrice.Value()
	}
	if g.Percentage.IsValueSet() {
		structMap["percentage"] = g.Percentage.Value()
	}
	return structMap
}

func (g *GetPricingSchemeResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Price         types.Optional[int]                       `json:"price"`
		SchemeType    types.Optional[string]                    `json:"scheme_type"`
		PriceBrackets types.Optional[[]GetPriceBracketResponse] `json:"price_brackets"`
		MinimumPrice  types.Optional[int]                       `json:"minimum_price"`
		Percentage    types.Optional[float64]                   `json:"percentage"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Price = temp.Price
	g.SchemeType = temp.SchemeType
	g.PriceBrackets = temp.PriceBrackets
	g.MinimumPrice = temp.MinimumPrice
	g.Percentage = temp.Percentage
	return nil
}
