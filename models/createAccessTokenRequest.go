package models

import (
	"encoding/json"
)

// Request for creating a new Access Token
type CreateAccessTokenRequest struct {
	ExpiresIn *int `json:"expires_in,omitempty"`
}

func (c *CreateAccessTokenRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateAccessTokenRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.ExpiresIn != nil {
		structMap["expires_in"] = c.ExpiresIn
	}
	return structMap
}

func (c *CreateAccessTokenRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ExpiresIn *int `json:"expires_in,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.ExpiresIn = temp.ExpiresIn
	return nil
}
