package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

type GetAutomaticAnticipationResponse struct {
	Enabled          types.Optional[bool]   `json:"enabled"`
	Type             types.Optional[string] `json:"type"`
	VolumePercentage types.Optional[int]    `json:"volume_percentage"`
	Delay            types.Optional[int]    `json:"delay"`
	Days             types.Optional[[]int]  `json:"days"`
}

func (g *GetAutomaticAnticipationResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetAutomaticAnticipationResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Enabled.IsValueSet() {
		structMap["enabled"] = g.Enabled.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.VolumePercentage.IsValueSet() {
		structMap["volume_percentage"] = g.VolumePercentage.Value()
	}
	if g.Delay.IsValueSet() {
		structMap["delay"] = g.Delay.Value()
	}
	if g.Days.IsValueSet() {
		structMap["days"] = g.Days.Value()
	}
	return structMap
}

func (g *GetAutomaticAnticipationResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Enabled          types.Optional[bool]   `json:"enabled"`
		Type             types.Optional[string] `json:"type"`
		VolumePercentage types.Optional[int]    `json:"volume_percentage"`
		Delay            types.Optional[int]    `json:"delay"`
		Days             types.Optional[[]int]  `json:"days"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Enabled = temp.Enabled
	g.Type = temp.Type
	g.VolumePercentage = temp.VolumePercentage
	g.Delay = temp.Delay
	g.Days = temp.Days
	return nil
}
