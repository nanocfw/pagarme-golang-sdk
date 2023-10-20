package models

import (
	"encoding/json"
)

// Atualização do valor mínimo da assinatura
type UpdateSubscriptionMinimumPriceRequest struct {
	MinimumPrice *int `json:"minimum_price,omitempty"`
}

func (u *UpdateSubscriptionMinimumPriceRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateSubscriptionMinimumPriceRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.MinimumPrice != nil {
		structMap["minimum_price"] = u.MinimumPrice
	}
	return structMap
}

func (u *UpdateSubscriptionMinimumPriceRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		MinimumPrice *int `json:"minimum_price,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.MinimumPrice = temp.MinimumPrice
	return nil
}
