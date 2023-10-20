package models

import (
	"encoding/json"
	"log"
	"time"
)

// Contains information to create a pix payment
type CreatePixPaymentRequest struct {
	ExpiresAt             *time.Time                  `json:"expires_at,omitempty"`
	ExpiresIn             *int                        `json:"expires_in,omitempty"`
	AdditionalInformation *[]PixAdditionalInformation `json:"additional_information,omitempty"`
}

func (c *CreatePixPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreatePixPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.ExpiresAt != nil {
		structMap["expires_at"] = c.ExpiresAt.Format(time.RFC3339)
	}
	if c.ExpiresIn != nil {
		structMap["expires_in"] = c.ExpiresIn
	}
	if c.AdditionalInformation != nil {
		structMap["additional_information"] = c.AdditionalInformation
	}
	return structMap
}

func (c *CreatePixPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ExpiresAt             *string                     `json:"expires_at,omitempty"`
		ExpiresIn             *int                        `json:"expires_in,omitempty"`
		AdditionalInformation *[]PixAdditionalInformation `json:"additional_information,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	if temp.ExpiresAt != nil {
		ExpiresAtVal, err := time.Parse(time.RFC3339, *temp.ExpiresAt)
		if err != nil {
			log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
		}
		c.ExpiresAt = &ExpiresAtVal
	}
	c.ExpiresIn = temp.ExpiresIn
	c.AdditionalInformation = temp.AdditionalInformation
	return nil
}
