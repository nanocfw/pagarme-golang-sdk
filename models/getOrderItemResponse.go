package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting an order item
type GetOrderItemResponse struct {
	Id          types.Optional[string] `json:"id"`
	Amount      types.Optional[int]    `json:"amount"`
	Description types.Optional[string] `json:"description"`
	Quantity    types.Optional[int]    `json:"quantity"`
	Category    types.Optional[string] `json:"category"`
	Code        types.Optional[string] `json:"code"`
}

func (g *GetOrderItemResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetOrderItemResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	if g.Quantity.IsValueSet() {
		structMap["quantity"] = g.Quantity.Value()
	}
	if g.Category.IsValueSet() {
		structMap["category"] = g.Category.Value()
	}
	if g.Code.IsValueSet() {
		structMap["code"] = g.Code.Value()
	}
	return structMap
}

func (g *GetOrderItemResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id          types.Optional[string] `json:"id"`
		Amount      types.Optional[int]    `json:"amount"`
		Description types.Optional[string] `json:"description"`
		Quantity    types.Optional[int]    `json:"quantity"`
		Category    types.Optional[string] `json:"category"`
		Code        types.Optional[string] `json:"code"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Amount = temp.Amount
	g.Description = temp.Description
	g.Quantity = temp.Quantity
	g.Category = temp.Category
	g.Code = temp.Code
	return nil
}
