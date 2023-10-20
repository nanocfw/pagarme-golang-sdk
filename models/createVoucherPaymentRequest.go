package models

import (
	"encoding/json"
)

// The settings for creating a voucher payment
type CreateVoucherPaymentRequest struct {
	StatementDescriptor *string            `json:"statement_descriptor,omitempty"`
	CardId              *string            `json:"card_id,omitempty"`
	CardToken           *string            `json:"card_token,omitempty"`
	Card                *CreateCardRequest `json:"Card,omitempty"`
	RecurrencyCycle     *string            `json:"recurrency_cycle,omitempty"`
}

func (c *CreateVoucherPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateVoucherPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.StatementDescriptor != nil {
		structMap["statement_descriptor"] = c.StatementDescriptor
	}
	if c.CardId != nil {
		structMap["card_id"] = c.CardId
	}
	if c.CardToken != nil {
		structMap["card_token"] = c.CardToken
	}
	if c.Card != nil {
		structMap["Card"] = c.Card
	}
	if c.RecurrencyCycle != nil {
		structMap["recurrency_cycle"] = c.RecurrencyCycle
	}
	return structMap
}

func (c *CreateVoucherPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StatementDescriptor *string            `json:"statement_descriptor,omitempty"`
		CardId              *string            `json:"card_id,omitempty"`
		CardToken           *string            `json:"card_token,omitempty"`
		Card                *CreateCardRequest `json:"Card,omitempty"`
		RecurrencyCycle     *string            `json:"recurrency_cycle,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.StatementDescriptor = temp.StatementDescriptor
	c.CardId = temp.CardId
	c.CardToken = temp.CardToken
	c.Card = temp.Card
	c.RecurrencyCycle = temp.RecurrencyCycle
	return nil
}
