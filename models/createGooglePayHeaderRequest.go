package models

import (
	"encoding/json"
)

// The GooglePay header request
type CreateGooglePayHeaderRequest struct {
	EphemeralPublicKey string `json:"ephemeral_public_key"`
}

func (c *CreateGooglePayHeaderRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateGooglePayHeaderRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["ephemeral_public_key"] = c.EphemeralPublicKey
	return structMap
}

func (c *CreateGooglePayHeaderRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		EphemeralPublicKey string `json:"ephemeral_public_key"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.EphemeralPublicKey = temp.EphemeralPublicKey
	return nil
}
