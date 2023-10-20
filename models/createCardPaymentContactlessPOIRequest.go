package models

import (
	"encoding/json"
)

type CreateCardPaymentContactlessPOIRequest struct {
	SystemName    string `json:"system_name"`
	Model         string `json:"model"`
	Provider      string `json:"provider"`
	SerialNumber  string `json:"serial_number"`
	VersionNumber string `json:"version_number"`
}

func (c *CreateCardPaymentContactlessPOIRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCardPaymentContactlessPOIRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["system_name"] = c.SystemName
	structMap["model"] = c.Model
	structMap["provider"] = c.Provider
	structMap["serial_number"] = c.SerialNumber
	structMap["version_number"] = c.VersionNumber
	return structMap
}

func (c *CreateCardPaymentContactlessPOIRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SystemName    string `json:"system_name"`
		Model         string `json:"model"`
		Provider      string `json:"provider"`
		SerialNumber  string `json:"serial_number"`
		VersionNumber string `json:"version_number"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.SystemName = temp.SystemName
	c.Model = temp.Model
	c.Provider = temp.Provider
	c.SerialNumber = temp.SerialNumber
	c.VersionNumber = temp.VersionNumber
	return nil
}
