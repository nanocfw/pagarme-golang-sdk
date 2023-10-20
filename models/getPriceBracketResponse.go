package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting a price bracket
type GetPriceBracketResponse struct {
	StartQuantity types.Optional[int] `json:"start_quantity"`
	Price         types.Optional[int] `json:"price"`
	EndQuantity   types.Optional[int] `json:"end_quantity"`
	OveragePrice  types.Optional[int] `json:"overage_price"`
}

func (g *GetPriceBracketResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetPriceBracketResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.StartQuantity.IsValueSet() {
		structMap["start_quantity"] = g.StartQuantity.Value()
	}
	if g.Price.IsValueSet() {
		structMap["price"] = g.Price.Value()
	}
	if g.EndQuantity.IsValueSet() {
		structMap["end_quantity"] = g.EndQuantity.Value()
	}
	if g.OveragePrice.IsValueSet() {
		structMap["overage_price"] = g.OveragePrice.Value()
	}
	return structMap
}

func (g *GetPriceBracketResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StartQuantity types.Optional[int] `json:"start_quantity"`
		Price         types.Optional[int] `json:"price"`
		EndQuantity   types.Optional[int] `json:"end_quantity"`
		OveragePrice  types.Optional[int] `json:"overage_price"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.StartQuantity = temp.StartQuantity
	g.Price = temp.Price
	g.EndQuantity = temp.EndQuantity
	g.OveragePrice = temp.OveragePrice
	return nil
}
