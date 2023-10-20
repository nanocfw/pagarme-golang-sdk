package models

import (
	"encoding/json"
)

type CreateAutomaticAnticipationSettingsRequest struct {
	Enabled          bool   `json:"enabled"`
	Type             string `json:"type"`
	VolumePercentage int    `json:"volume_percentage"`
	Delay            int    `json:"delay"`
	Days             []int  `json:"days"`
}

func (c *CreateAutomaticAnticipationSettingsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateAutomaticAnticipationSettingsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["enabled"] = c.Enabled
	structMap["type"] = c.Type
	structMap["volume_percentage"] = c.VolumePercentage
	structMap["delay"] = c.Delay
	structMap["days"] = c.Days
	return structMap
}

func (c *CreateAutomaticAnticipationSettingsRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Enabled          bool   `json:"enabled"`
		Type             string `json:"type"`
		VolumePercentage int    `json:"volume_percentage"`
		Delay            int    `json:"delay"`
		Days             []int  `json:"days"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Enabled = temp.Enabled
	c.Type = temp.Type
	c.VolumePercentage = temp.VolumePercentage
	c.Delay = temp.Delay
	c.Days = temp.Days
	return nil
}
