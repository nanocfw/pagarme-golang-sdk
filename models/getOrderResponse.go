package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for getting an Order
type GetOrderResponse struct {
	Id         types.Optional[string]                       `json:"id"`
	Code       types.Optional[string]                       `json:"code"`
	Currency   types.Optional[string]                       `json:"currency"`
	Items      types.Optional[[]GetOrderItemResponse]       `json:"items"`
	Customer   types.Optional[GetCustomerResponse]          `json:"customer"`
	Status     types.Optional[string]                       `json:"status"`
	CreatedAt  types.Optional[time.Time]                    `json:"created_at"`
	UpdatedAt  types.Optional[time.Time]                    `json:"updated_at"`
	Charges    types.Optional[[]GetChargeResponse]          `json:"charges"`
	InvoiceUrl types.Optional[string]                       `json:"invoice_url"`
	Shipping   types.Optional[GetShippingResponse]          `json:"shipping"`
	Metadata   types.Optional[map[string]string]            `json:"metadata"`
	Checkouts  types.Optional[[]GetCheckoutPaymentResponse] `json:"checkouts"`
	Ip         types.Optional[string]                       `json:"ip"`
	SessionId  types.Optional[string]                       `json:"session_id"`
	Location   types.Optional[GetLocationResponse]          `json:"location"`
	Device     types.Optional[GetDeviceResponse]            `json:"device"`
	Closed     types.Optional[bool]                         `json:"closed"`
}

func (g *GetOrderResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetOrderResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Code.IsValueSet() {
		structMap["code"] = g.Code.Value()
	}
	if g.Currency.IsValueSet() {
		structMap["currency"] = g.Currency.Value()
	}
	if g.Items.IsValueSet() {
		structMap["items"] = g.Items.Value()
	}
	if g.Customer.IsValueSet() {
		structMap["customer"] = g.Customer.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.CreatedAt.IsValueSet() {
		var CreatedAtVal *string = nil
		if g.CreatedAt.Value() != nil {
			val := g.CreatedAt.Value().Format(time.RFC3339)
			CreatedAtVal = &val
		}
		structMap["created_at"] = CreatedAtVal
	}
	if g.UpdatedAt.IsValueSet() {
		var UpdatedAtVal *string = nil
		if g.UpdatedAt.Value() != nil {
			val := g.UpdatedAt.Value().Format(time.RFC3339)
			UpdatedAtVal = &val
		}
		structMap["updated_at"] = UpdatedAtVal
	}
	if g.Charges.IsValueSet() {
		structMap["charges"] = g.Charges.Value()
	}
	if g.InvoiceUrl.IsValueSet() {
		structMap["invoice_url"] = g.InvoiceUrl.Value()
	}
	if g.Shipping.IsValueSet() {
		structMap["shipping"] = g.Shipping.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.Checkouts.IsValueSet() {
		structMap["checkouts"] = g.Checkouts.Value()
	}
	if g.Ip.IsValueSet() {
		structMap["ip"] = g.Ip.Value()
	}
	if g.SessionId.IsValueSet() {
		structMap["session_id"] = g.SessionId.Value()
	}
	if g.Location.IsValueSet() {
		structMap["location"] = g.Location.Value()
	}
	if g.Device.IsValueSet() {
		structMap["device"] = g.Device.Value()
	}
	if g.Closed.IsValueSet() {
		structMap["closed"] = g.Closed.Value()
	}
	return structMap
}

func (g *GetOrderResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id         types.Optional[string]                       `json:"id"`
		Code       types.Optional[string]                       `json:"code"`
		Currency   types.Optional[string]                       `json:"currency"`
		Items      types.Optional[[]GetOrderItemResponse]       `json:"items"`
		Customer   types.Optional[GetCustomerResponse]          `json:"customer"`
		Status     types.Optional[string]                       `json:"status"`
		CreatedAt  types.Optional[string]                       `json:"created_at"`
		UpdatedAt  types.Optional[string]                       `json:"updated_at"`
		Charges    types.Optional[[]GetChargeResponse]          `json:"charges"`
		InvoiceUrl types.Optional[string]                       `json:"invoice_url"`
		Shipping   types.Optional[GetShippingResponse]          `json:"shipping"`
		Metadata   types.Optional[map[string]string]            `json:"metadata"`
		Checkouts  types.Optional[[]GetCheckoutPaymentResponse] `json:"checkouts"`
		Ip         types.Optional[string]                       `json:"ip"`
		SessionId  types.Optional[string]                       `json:"session_id"`
		Location   types.Optional[GetLocationResponse]          `json:"location"`
		Device     types.Optional[GetDeviceResponse]            `json:"device"`
		Closed     types.Optional[bool]                         `json:"closed"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Code = temp.Code
	g.Currency = temp.Currency
	g.Items = temp.Items
	g.Customer = temp.Customer
	g.Status = temp.Status
	g.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
	if temp.CreatedAt.Value() != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		g.CreatedAt.SetValue(&CreatedAtVal)
	}
	g.UpdatedAt.ShouldSetValue(temp.UpdatedAt.IsValueSet())
	if temp.UpdatedAt.Value() != nil {
		UpdatedAtVal, err := time.Parse(time.RFC3339, (*temp.UpdatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
		}
		g.UpdatedAt.SetValue(&UpdatedAtVal)
	}
	g.Charges = temp.Charges
	g.InvoiceUrl = temp.InvoiceUrl
	g.Shipping = temp.Shipping
	g.Metadata = temp.Metadata
	g.Checkouts = temp.Checkouts
	g.Ip = temp.Ip
	g.SessionId = temp.SessionId
	g.Location = temp.Location
	g.Device = temp.Device
	g.Closed = temp.Closed
	return nil
}
