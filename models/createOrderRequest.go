package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Request for creating an order
type CreateOrderRequest struct {
	Items            []CreateOrderItemRequest          `json:"items"`
	Customer         CreateCustomerRequest             `json:"customer"`
	Payments         []CreatePaymentRequest            `json:"payments"`
	Code             string                            `json:"code"`
	CustomerId       types.Optional[string]            `json:"customer_id"`
	Shipping         *CreateShippingRequest            `json:"shipping,omitempty"`
	Metadata         types.Optional[map[string]string] `json:"metadata"`
	AntifraudEnabled *bool                             `json:"antifraud_enabled,omitempty"`
	Ip               *string                           `json:"ip,omitempty"`
	SessionId        *string                           `json:"session_id,omitempty"`
	Location         *CreateLocationRequest            `json:"location,omitempty"`
	Device           *CreateDeviceRequest              `json:"device,omitempty"`
	Closed           bool                              `json:"closed"`
	Currency         *string                           `json:"currency,omitempty"`
	Antifraud        *CreateAntifraudRequest           `json:"antifraud,omitempty"`
	Submerchant      *CreateSubMerchantRequest         `json:"submerchant,omitempty"`
}

func (c *CreateOrderRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateOrderRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["items"] = c.Items
	structMap["customer"] = c.Customer
	structMap["payments"] = c.Payments
	structMap["code"] = c.Code
	if c.CustomerId.IsValueSet() {
		structMap["customer_id"] = c.CustomerId.Value()
	}
	if c.Shipping != nil {
		structMap["shipping"] = c.Shipping
	}
	if c.Metadata.IsValueSet() {
		structMap["metadata"] = c.Metadata.Value()
	}
	if c.AntifraudEnabled != nil {
		structMap["antifraud_enabled"] = c.AntifraudEnabled
	}
	if c.Ip != nil {
		structMap["ip"] = c.Ip
	}
	if c.SessionId != nil {
		structMap["session_id"] = c.SessionId
	}
	if c.Location != nil {
		structMap["location"] = c.Location
	}
	if c.Device != nil {
		structMap["device"] = c.Device
	}
	structMap["closed"] = c.Closed
	if c.Currency != nil {
		structMap["currency"] = c.Currency
	}
	if c.Antifraud != nil {
		structMap["antifraud"] = c.Antifraud
	}
	if c.Submerchant != nil {
		structMap["submerchant"] = c.Submerchant
	}
	return structMap
}

func (c *CreateOrderRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Items            []CreateOrderItemRequest          `json:"items"`
		Customer         CreateCustomerRequest             `json:"customer"`
		Payments         []CreatePaymentRequest            `json:"payments"`
		Code             string                            `json:"code"`
		CustomerId       types.Optional[string]            `json:"customer_id"`
		Shipping         *CreateShippingRequest            `json:"shipping,omitempty"`
		Metadata         types.Optional[map[string]string] `json:"metadata"`
		AntifraudEnabled *bool                             `json:"antifraud_enabled,omitempty"`
		Ip               *string                           `json:"ip,omitempty"`
		SessionId        *string                           `json:"session_id,omitempty"`
		Location         *CreateLocationRequest            `json:"location,omitempty"`
		Device           *CreateDeviceRequest              `json:"device,omitempty"`
		Closed           bool                              `json:"closed"`
		Currency         *string                           `json:"currency,omitempty"`
		Antifraud        *CreateAntifraudRequest           `json:"antifraud,omitempty"`
		Submerchant      *CreateSubMerchantRequest         `json:"submerchant,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Items = temp.Items
	c.Customer = temp.Customer
	c.Payments = temp.Payments
	c.Code = temp.Code
	c.CustomerId = temp.CustomerId
	c.Shipping = temp.Shipping
	c.Metadata = temp.Metadata
	c.AntifraudEnabled = temp.AntifraudEnabled
	c.Ip = temp.Ip
	c.SessionId = temp.SessionId
	c.Location = temp.Location
	c.Device = temp.Device
	c.Closed = temp.Closed
	c.Currency = temp.Currency
	c.Antifraud = temp.Antifraud
	c.Submerchant = temp.Submerchant
	return nil
}
