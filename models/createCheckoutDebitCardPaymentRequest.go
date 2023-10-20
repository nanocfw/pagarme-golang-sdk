package models

import (
	"encoding/json"
)

// Checkout credit card payment request
type CreateCheckoutDebitCardPaymentRequest struct {
	StatementDescriptor *string                            `json:"statement_descriptor,omitempty"`
	Authentication      CreatePaymentAuthenticationRequest `json:"authentication"`
}

func (c *CreateCheckoutDebitCardPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCheckoutDebitCardPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.StatementDescriptor != nil {
		structMap["statement_descriptor"] = c.StatementDescriptor
	}
	structMap["authentication"] = c.Authentication
	return structMap
}

func (c *CreateCheckoutDebitCardPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StatementDescriptor *string                            `json:"statement_descriptor,omitempty"`
		Authentication      CreatePaymentAuthenticationRequest `json:"authentication"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.StatementDescriptor = temp.StatementDescriptor
	c.Authentication = temp.Authentication
	return nil
}
