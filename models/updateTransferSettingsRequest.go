package models

import (
	"encoding/json"
)

type UpdateTransferSettingsRequest struct {
	TransferEnabled  string `json:"transfer_enabled"`
	TransferInterval string `json:"transfer_interval"`
	TransferDay      string `json:"transfer_day"`
}

func (u *UpdateTransferSettingsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateTransferSettingsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["transfer_enabled"] = u.TransferEnabled
	structMap["transfer_interval"] = u.TransferInterval
	structMap["transfer_day"] = u.TransferDay
	return structMap
}

func (u *UpdateTransferSettingsRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		TransferEnabled  string `json:"transfer_enabled"`
		TransferInterval string `json:"transfer_interval"`
		TransferDay      string `json:"transfer_day"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.TransferEnabled = temp.TransferEnabled
	u.TransferInterval = temp.TransferInterval
	u.TransferDay = temp.TransferDay
	return nil
}
