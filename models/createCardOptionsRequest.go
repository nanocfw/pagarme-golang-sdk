package models

import (
	"encoding/json"
)

// Options for creating the card
type CreateCardOptionsRequest struct {
	VerifyCard bool `json:"verify_card"`
}

func (c *CreateCardOptionsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCardOptionsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["verify_card"] = c.VerifyCard
	return structMap
}

func (c *CreateCardOptionsRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		VerifyCard bool `json:"verify_card"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.VerifyCard = temp.VerifyCard
	return nil
}
