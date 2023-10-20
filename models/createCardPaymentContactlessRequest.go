package models

import (
	"encoding/json"
)

// The card payment contactless request
type CreateCardPaymentContactlessRequest struct {
	Type      string                   `json:"type"`
	ApplePay  *CreateApplePayRequest   `json:"apple_pay,omitempty"`
	GooglePay *CreateGooglePayRequest  `json:"google_pay,omitempty"`
	Emv       *CreateEmvDecryptRequest `json:"emv,omitempty"`
}

func (c *CreateCardPaymentContactlessRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCardPaymentContactlessRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["type"] = c.Type
	if c.ApplePay != nil {
		structMap["apple_pay"] = c.ApplePay
	}
	if c.GooglePay != nil {
		structMap["google_pay"] = c.GooglePay
	}
	if c.Emv != nil {
		structMap["emv"] = c.Emv
	}
	return structMap
}

func (c *CreateCardPaymentContactlessRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Type      string                   `json:"type"`
		ApplePay  *CreateApplePayRequest   `json:"apple_pay,omitempty"`
		GooglePay *CreateGooglePayRequest  `json:"google_pay,omitempty"`
		Emv       *CreateEmvDecryptRequest `json:"emv,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Type = temp.Type
	c.ApplePay = temp.ApplePay
	c.GooglePay = temp.GooglePay
	c.Emv = temp.Emv
	return nil
}
