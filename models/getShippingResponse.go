package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting the shipping data
type GetShippingResponse struct {
	Amount                types.Optional[int]                `json:"amount"`
	Description           types.Optional[string]             `json:"description"`
	RecipientName         types.Optional[string]             `json:"recipient_name"`
	RecipientPhone        types.Optional[string]             `json:"recipient_phone"`
	Address               types.Optional[GetAddressResponse] `json:"address"`
	MaxDeliveryDate       types.Optional[time.Time]          `json:"max_delivery_date"`
	EstimatedDeliveryDate types.Optional[time.Time]          `json:"estimated_delivery_date"`
	Type                  types.Optional[string]             `json:"type"`
}

func (g *GetShippingResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetShippingResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	if g.RecipientName.IsValueSet() {
		structMap["recipient_name"] = g.RecipientName.Value()
	}
	if g.RecipientPhone.IsValueSet() {
		structMap["recipient_phone"] = g.RecipientPhone.Value()
	}
	if g.Address.IsValueSet() {
		structMap["address"] = g.Address.Value()
	}
	if g.MaxDeliveryDate.IsValueSet() {
		var MaxDeliveryDateVal *string = nil
		if g.MaxDeliveryDate.Value() != nil {
			val := g.MaxDeliveryDate.Value().Format(time.RFC3339)
			MaxDeliveryDateVal = &val
		}
		structMap["max_delivery_date"] = MaxDeliveryDateVal
	}
	if g.EstimatedDeliveryDate.IsValueSet() {
		var EstimatedDeliveryDateVal *string = nil
		if g.EstimatedDeliveryDate.Value() != nil {
			val := g.EstimatedDeliveryDate.Value().Format(time.RFC3339)
			EstimatedDeliveryDateVal = &val
		}
		structMap["estimated_delivery_date"] = EstimatedDeliveryDateVal
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	return structMap
}

func (g *GetShippingResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount                types.Optional[int]                `json:"amount"`
		Description           types.Optional[string]             `json:"description"`
		RecipientName         types.Optional[string]             `json:"recipient_name"`
		RecipientPhone        types.Optional[string]             `json:"recipient_phone"`
		Address               types.Optional[GetAddressResponse] `json:"address"`
		MaxDeliveryDate       types.Optional[string]             `json:"max_delivery_date"`
		EstimatedDeliveryDate types.Optional[string]             `json:"estimated_delivery_date"`
		Type                  types.Optional[string]             `json:"type"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Amount = temp.Amount
	g.Description = temp.Description
	g.RecipientName = temp.RecipientName
	g.RecipientPhone = temp.RecipientPhone
	g.Address = temp.Address
	g.MaxDeliveryDate.ShouldSetValue(temp.MaxDeliveryDate.IsValueSet())
	if temp.MaxDeliveryDate.Value() != nil {
		MaxDeliveryDateVal, err := time.Parse(time.RFC3339, (*temp.MaxDeliveryDate.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse max_delivery_date as % s format.", time.RFC3339)
		}
		g.MaxDeliveryDate.SetValue(&MaxDeliveryDateVal)
	}
	g.EstimatedDeliveryDate.ShouldSetValue(temp.EstimatedDeliveryDate.IsValueSet())
	if temp.EstimatedDeliveryDate.Value() != nil {
		EstimatedDeliveryDateVal, err := time.Parse(time.RFC3339, (*temp.EstimatedDeliveryDate.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse estimated_delivery_date as % s format.", time.RFC3339)
		}
		g.EstimatedDeliveryDate.SetValue(&EstimatedDeliveryDateVal)
	}
	g.Type = temp.Type
	return nil
}
