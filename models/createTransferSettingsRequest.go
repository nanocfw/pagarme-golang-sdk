package models

import (
	"encoding/json"
)

// Informações de transferência do recebedor
type CreateTransferSettingsRequest struct {
	TransferEnabled  bool   `json:"transfer_enabled"`
	TransferInterval string `json:"transfer_interval"`
	TransferDay      int    `json:"transfer_day"`
}

func (c *CreateTransferSettingsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateTransferSettingsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["transfer_enabled"] = c.TransferEnabled
	structMap["transfer_interval"] = c.TransferInterval
	structMap["transfer_day"] = c.TransferDay
	return structMap
}

func (c *CreateTransferSettingsRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		TransferEnabled  bool   `json:"transfer_enabled"`
		TransferInterval string `json:"transfer_interval"`
		TransferDay      int    `json:"transfer_day"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.TransferEnabled = temp.TransferEnabled
	c.TransferInterval = temp.TransferInterval
	c.TransferDay = temp.TransferDay
	return nil
}
