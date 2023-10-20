package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Request for creating a new charge
type CreateChargeRequest struct {
	Code       types.Optional[string]                 `json:"code"`
	Amount     int                                    `json:"amount"`
	CustomerId types.Optional[string]                 `json:"customer_id"`
	Customer   types.Optional[CreateCustomerRequest]  `json:"customer"`
	Payment    CreatePaymentRequest                   `json:"payment"`
	Metadata   types.Optional[map[string]string]      `json:"metadata"`
	DueAt      types.Optional[time.Time]              `json:"due_at"`
	Antifraud  types.Optional[CreateAntifraudRequest] `json:"antifraud"`
	OrderId    string                                 `json:"order_id"`
}

func (c *CreateChargeRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateChargeRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Code.IsValueSet() {
		structMap["code"] = c.Code.Value()
	}
	structMap["amount"] = c.Amount
	if c.CustomerId.IsValueSet() {
		structMap["customer_id"] = c.CustomerId.Value()
	}
	if c.Customer.IsValueSet() {
		structMap["customer"] = c.Customer.Value()
	}
	structMap["payment"] = c.Payment
	if c.Metadata.IsValueSet() {
		structMap["metadata"] = c.Metadata.Value()
	}
	if c.DueAt.IsValueSet() {
		var DueAtVal *string = nil
		if c.DueAt.Value() != nil {
			val := c.DueAt.Value().Format(time.RFC3339)
			DueAtVal = &val
		}
		structMap["due_at"] = DueAtVal
	}
	if c.Antifraud.IsValueSet() {
		structMap["antifraud"] = c.Antifraud.Value()
	}
	structMap["order_id"] = c.OrderId
	return structMap
}

func (c *CreateChargeRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Code       types.Optional[string]                 `json:"code"`
		Amount     int                                    `json:"amount"`
		CustomerId types.Optional[string]                 `json:"customer_id"`
		Customer   types.Optional[CreateCustomerRequest]  `json:"customer"`
		Payment    CreatePaymentRequest                   `json:"payment"`
		Metadata   types.Optional[map[string]string]      `json:"metadata"`
		DueAt      types.Optional[string]                 `json:"due_at"`
		Antifraud  types.Optional[CreateAntifraudRequest] `json:"antifraud"`
		OrderId    string                                 `json:"order_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Code = temp.Code
	c.Amount = temp.Amount
	c.CustomerId = temp.CustomerId
	c.Customer = temp.Customer
	c.Payment = temp.Payment
	c.Metadata = temp.Metadata
	c.DueAt.ShouldSetValue(temp.DueAt.IsValueSet())
	if temp.DueAt.Value() != nil {
		DueAtVal, err := time.Parse(time.RFC3339, (*temp.DueAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
		}
		c.DueAt.SetValue(&DueAtVal)
	}
	c.Antifraud = temp.Antifraud
	c.OrderId = temp.OrderId
	return nil
}
