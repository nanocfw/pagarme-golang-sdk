package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/apimatic/go-core-runtime/types"
)

// Contains the settings for creating a boleto payment
type CreateBoletoPaymentRequest struct {
	Retries             int                                   `json:"retries"`
	Bank                types.Optional[string]                `json:"bank"`
	Instructions        string                                `json:"instructions"`
	DueAt               types.Optional[time.Time]             `json:"due_at"`
	BillingAddress      CreateAddressRequest                  `json:"billing_address"`
	BillingAddressId    types.Optional[string]                `json:"billing_address_id"`
	NossoNumero         types.Optional[string]                `json:"nosso_numero"`
	DocumentNumber      string                                `json:"document_number"`
	StatementDescriptor string                                `json:"statement_descriptor"`
	Interest            types.Optional[CreateInterestRequest] `json:"interest"`
	Fine                types.Optional[CreateFineRequest]     `json:"fine"`
	MaxDaysToPayPastDue types.Optional[int]                   `json:"max_days_to_pay_past_due"`
}

func (c *CreateBoletoPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateBoletoPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["retries"] = c.Retries
	if c.Bank.IsValueSet() {
		structMap["bank"] = c.Bank.Value()
	}
	structMap["instructions"] = c.Instructions
	if c.DueAt.IsValueSet() {
		var DueAtVal *string = nil
		if c.DueAt.Value() != nil {
			val := c.DueAt.Value().Format(time.RFC3339)
			DueAtVal = &val
		}
		structMap["due_at"] = DueAtVal
	}
	structMap["billing_address"] = c.BillingAddress
	if c.BillingAddressId.IsValueSet() {
		structMap["billing_address_id"] = c.BillingAddressId.Value()
	}
	if c.NossoNumero.IsValueSet() {
		structMap["nosso_numero"] = c.NossoNumero.Value()
	}
	structMap["document_number"] = c.DocumentNumber
	structMap["statement_descriptor"] = c.StatementDescriptor
	if c.Interest.IsValueSet() {
		structMap["interest"] = c.Interest.Value()
	}
	if c.Fine.IsValueSet() {
		structMap["fine"] = c.Fine.Value()
	}
	if c.MaxDaysToPayPastDue.IsValueSet() {
		structMap["max_days_to_pay_past_due"] = c.MaxDaysToPayPastDue.Value()
	}
	return structMap
}

func (c *CreateBoletoPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Retries             int                                   `json:"retries"`
		Bank                types.Optional[string]                `json:"bank"`
		Instructions        string                                `json:"instructions"`
		DueAt               types.Optional[string]                `json:"due_at"`
		BillingAddress      CreateAddressRequest                  `json:"billing_address"`
		BillingAddressId    types.Optional[string]                `json:"billing_address_id"`
		NossoNumero         types.Optional[string]                `json:"nosso_numero"`
		DocumentNumber      string                                `json:"document_number"`
		StatementDescriptor string                                `json:"statement_descriptor"`
		Interest            types.Optional[CreateInterestRequest] `json:"interest"`
		Fine                types.Optional[CreateFineRequest]     `json:"fine"`
		MaxDaysToPayPastDue types.Optional[int]                   `json:"max_days_to_pay_past_due"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Retries = temp.Retries
	c.Bank = temp.Bank
	c.Instructions = temp.Instructions
	c.DueAt.ShouldSetValue(temp.DueAt.IsValueSet())
	if temp.DueAt.Value() != nil {
		DueAtVal, err := time.Parse(time.RFC3339, (*temp.DueAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
		}
		c.DueAt.SetValue(&DueAtVal)
	}
	c.BillingAddress = temp.BillingAddress
	c.BillingAddressId = temp.BillingAddressId
	c.NossoNumero = temp.NossoNumero
	c.DocumentNumber = temp.DocumentNumber
	c.StatementDescriptor = temp.StatementDescriptor
	c.Interest = temp.Interest
	c.Fine = temp.Fine
	c.MaxDaysToPayPastDue = temp.MaxDaysToPayPastDue
	return nil
}
