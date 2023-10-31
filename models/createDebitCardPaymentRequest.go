package models

import (
	"encoding/json"
)

// The settings for creating a debit card payment
type CreateDebitCardPaymentRequest struct {
	StatementDescriptor *string                              `json:"statement_descriptor,omitempty"`
	Card                *CreateCardRequest                   `json:"card,omitempty"`
	CardId              *string                              `json:"card_id,omitempty"`
	CardToken           *string                              `json:"card_token,omitempty"`
	Recurrence          *bool                                `json:"recurrence,omitempty"`
	Capture             *bool                                `json:"capture,omitempty"`
	Authentication      *CreatePaymentAuthenticationRequest  `json:"authentication,omitempty"`
	Token               *CreateCardPaymentContactlessRequest `json:"token,omitempty"`
}

func (c *CreateDebitCardPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateDebitCardPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.StatementDescriptor != nil {
		structMap["statement_descriptor"] = c.StatementDescriptor
	}
	if c.Card != nil {
		structMap["card"] = c.Card
	}
	if c.CardId != nil {
		structMap["card_id"] = c.CardId
	}
	if c.CardToken != nil {
		structMap["card_token"] = c.CardToken
	}
	if c.Recurrence != nil {
		structMap["recurrence"] = c.Recurrence
	}
	if c.Capture != nil {
		structMap["capture"] = c.Capture
	}
	if c.Authentication != nil {
		structMap["authentication"] = c.Authentication
	}
	if c.Token != nil {
		structMap["token"] = c.Token
	}
	return structMap
}

func (c *CreateDebitCardPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StatementDescriptor *string                              `json:"statement_descriptor,omitempty"`
		Card                *CreateCardRequest                   `json:"card,omitempty"`
		CardId              *string                              `json:"card_id,omitempty"`
		CardToken           *string                              `json:"card_token,omitempty"`
		Recurrence          *bool                                `json:"recurrence,omitempty"`
		Capture             *bool                                `json:"capture,omitempty"`
		Authentication      *CreatePaymentAuthenticationRequest  `json:"authentication,omitempty"`
		Token               *CreateCardPaymentContactlessRequest `json:"token,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.StatementDescriptor = temp.StatementDescriptor
	c.Card = temp.Card
	c.CardId = temp.CardId
	c.CardToken = temp.CardToken
	c.Recurrence = temp.Recurrence
	c.Capture = temp.Capture
	c.Authentication = temp.Authentication
	c.Token = temp.Token
	return nil
}
