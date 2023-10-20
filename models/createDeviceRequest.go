package models

import (
	"encoding/json"
)

// Request for creating a device
type CreateDeviceRequest struct {
	Platform *string `json:"platform,omitempty"`
}

func (c *CreateDeviceRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateDeviceRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Platform != nil {
		structMap["platform"] = c.Platform
	}
	return structMap
}

func (c *CreateDeviceRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Platform *string `json:"platform,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Platform = temp.Platform
	return nil
}
