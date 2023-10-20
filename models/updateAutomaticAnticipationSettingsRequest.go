package models

import (
	"encoding/json"
)

type UpdateAutomaticAnticipationSettingsRequest struct {
	Enabled          *bool   `json:"enabled,omitempty"`
	Type             *string `json:"type,omitempty"`
	VolumePercentage *int    `json:"volume_percentage,omitempty"`
	Delay            *int    `json:"delay,omitempty"`
	Days             *int    `json:"days,omitempty"`
}

func (u *UpdateAutomaticAnticipationSettingsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateAutomaticAnticipationSettingsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.Enabled != nil {
		structMap["enabled"] = u.Enabled
	}
	if u.Type != nil {
		structMap["type"] = u.Type
	}
	if u.VolumePercentage != nil {
		structMap["volume_percentage"] = u.VolumePercentage
	}
	if u.Delay != nil {
		structMap["delay"] = u.Delay
	}
	if u.Days != nil {
		structMap["days"] = u.Days
	}
	return structMap
}

func (u *UpdateAutomaticAnticipationSettingsRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Enabled          *bool   `json:"enabled,omitempty"`
		Type             *string `json:"type,omitempty"`
		VolumePercentage *int    `json:"volume_percentage,omitempty"`
		Delay            *int    `json:"delay,omitempty"`
		Days             *int    `json:"days,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Enabled = temp.Enabled
	u.Type = temp.Type
	u.VolumePercentage = temp.VolumePercentage
	u.Delay = temp.Delay
	u.Days = temp.Days
	return nil
}
