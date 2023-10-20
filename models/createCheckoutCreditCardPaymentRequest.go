package models

import (
	"encoding/json"
)

// Checkout card payment request
type CreateCheckoutCreditCardPaymentRequest struct {
	StatementDescriptor *string                                       `json:"statement_descriptor,omitempty"`
	Installments        *[]CreateCheckoutCardInstallmentOptionRequest `json:"installments,omitempty"`
	Authentication      *CreatePaymentAuthenticationRequest           `json:"authentication,omitempty"`
	Capture             *bool                                         `json:"capture,omitempty"`
}

func (c *CreateCheckoutCreditCardPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCheckoutCreditCardPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.StatementDescriptor != nil {
		structMap["statement_descriptor"] = c.StatementDescriptor
	}
	if c.Installments != nil {
		structMap["installments"] = c.Installments
	}
	if c.Authentication != nil {
		structMap["authentication"] = c.Authentication
	}
	if c.Capture != nil {
		structMap["capture"] = c.Capture
	}
	return structMap
}

func (c *CreateCheckoutCreditCardPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StatementDescriptor *string                                       `json:"statement_descriptor,omitempty"`
		Installments        *[]CreateCheckoutCardInstallmentOptionRequest `json:"installments,omitempty"`
		Authentication      *CreatePaymentAuthenticationRequest           `json:"authentication,omitempty"`
		Capture             *bool                                         `json:"capture,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.StatementDescriptor = temp.StatementDescriptor
	c.Installments = temp.Installments
	c.Authentication = temp.Authentication
	c.Capture = temp.Capture
	return nil
}
