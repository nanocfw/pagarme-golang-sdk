package models

import (
	"encoding/json"
)

// Request for updating the card from a subscription
type UpdateSubscriptionCardRequest struct {
	Card   CreateCardRequest `json:"card"`
	CardId string            `json:"card_id"`
}

func (u *UpdateSubscriptionCardRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateSubscriptionCardRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["card"] = u.Card
	structMap["card_id"] = u.CardId
	return structMap
}

func (u *UpdateSubscriptionCardRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Card   CreateCardRequest `json:"card"`
		CardId string            `json:"card_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Card = temp.Card
	u.CardId = temp.CardId
	return nil
}
