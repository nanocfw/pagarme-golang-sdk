package models

import (
	"encoding/json"
)

// Request for updating card data
type UpdateChargeCardRequest struct {
	UpdateSubscription bool              `json:"update_subscription"`
	CardId             string            `json:"card_id"`
	Card               CreateCardRequest `json:"card"`
	Recurrence         bool              `json:"recurrence"`
}

func (u *UpdateChargeCardRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateChargeCardRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["update_subscription"] = u.UpdateSubscription
	structMap["card_id"] = u.CardId
	structMap["card"] = u.Card
	structMap["recurrence"] = u.Recurrence
	return structMap
}

func (u *UpdateChargeCardRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		UpdateSubscription bool              `json:"update_subscription"`
		CardId             string            `json:"card_id"`
		Card               CreateCardRequest `json:"card"`
		Recurrence         bool              `json:"recurrence"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.UpdateSubscription = temp.UpdateSubscription
	u.CardId = temp.CardId
	u.Card = temp.Card
	u.Recurrence = temp.Recurrence
	return nil
}
