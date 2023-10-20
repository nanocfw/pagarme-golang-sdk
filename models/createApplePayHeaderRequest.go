package models

import (
	"encoding/json"
)

// The ApplePay header request
type CreateApplePayHeaderRequest struct {
	PublicKeyHash      *string `json:"public_key_hash,omitempty"`
	EphemeralPublicKey string  `json:"ephemeral_public_key"`
	TransactionId      *string `json:"transaction_id,omitempty"`
}

func (c *CreateApplePayHeaderRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateApplePayHeaderRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.PublicKeyHash != nil {
		structMap["public_key_hash"] = c.PublicKeyHash
	}
	structMap["ephemeral_public_key"] = c.EphemeralPublicKey
	if c.TransactionId != nil {
		structMap["transaction_id"] = c.TransactionId
	}
	return structMap
}

func (c *CreateApplePayHeaderRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PublicKeyHash      *string `json:"public_key_hash,omitempty"`
		EphemeralPublicKey string  `json:"ephemeral_public_key"`
		TransactionId      *string `json:"transaction_id,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.PublicKeyHash = temp.PublicKeyHash
	c.EphemeralPublicKey = temp.EphemeralPublicKey
	c.TransactionId = temp.TransactionId
	return nil
}
