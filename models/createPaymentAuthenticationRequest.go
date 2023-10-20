package models

import (
	"encoding/json"
)

// The payment authentication request
type CreatePaymentAuthenticationRequest struct {
	Type         string                    `json:"type"`
	ThreedSecure CreateThreeDSecureRequest `json:"threed_secure"`
}

func (c *CreatePaymentAuthenticationRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreatePaymentAuthenticationRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["type"] = c.Type
	structMap["threed_secure"] = c.ThreedSecure
	return structMap
}

func (c *CreatePaymentAuthenticationRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Type         string                    `json:"type"`
		ThreedSecure CreateThreeDSecureRequest `json:"threed_secure"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Type = temp.Type
	c.ThreedSecure = temp.ThreedSecure
	return nil
}
