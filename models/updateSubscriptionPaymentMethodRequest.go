package models

import (
	"encoding/json"
)

// Request for updating a subscription's payment method
type UpdateSubscriptionPaymentMethodRequest struct {
	PaymentMethod string                           `json:"payment_method"`
	CardId        string                           `json:"card_id"`
	Card          CreateCardRequest                `json:"card"`
	CardToken     *string                          `json:"card_token,omitempty"`
	Boleto        *CreateSubscriptionBoletoRequest `json:"boleto,omitempty"`
}

func (u *UpdateSubscriptionPaymentMethodRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateSubscriptionPaymentMethodRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["payment_method"] = u.PaymentMethod
	structMap["card_id"] = u.CardId
	structMap["card"] = u.Card
	if u.CardToken != nil {
		structMap["card_token"] = u.CardToken
	}
	if u.Boleto != nil {
		structMap["boleto"] = u.Boleto
	}
	return structMap
}

func (u *UpdateSubscriptionPaymentMethodRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PaymentMethod string                           `json:"payment_method"`
		CardId        string                           `json:"card_id"`
		Card          CreateCardRequest                `json:"card"`
		CardToken     *string                          `json:"card_token,omitempty"`
		Boleto        *CreateSubscriptionBoletoRequest `json:"boleto,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.PaymentMethod = temp.PaymentMethod
	u.CardId = temp.CardId
	u.Card = temp.Card
	u.CardToken = temp.CardToken
	u.Boleto = temp.Boleto
	return nil
}
