package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

type GetTransferSettingsResponse struct {
	TransferEnabled  types.Optional[bool]   `json:"transfer_enabled"`
	TransferInterval types.Optional[string] `json:"transfer_interval"`
	TransferDay      types.Optional[int]    `json:"transfer_day"`
}

func (g *GetTransferSettingsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetTransferSettingsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.TransferEnabled.IsValueSet() {
		structMap["transfer_enabled"] = g.TransferEnabled.Value()
	}
	if g.TransferInterval.IsValueSet() {
		structMap["transfer_interval"] = g.TransferInterval.Value()
	}
	if g.TransferDay.IsValueSet() {
		structMap["transfer_day"] = g.TransferDay.Value()
	}
	return structMap
}

func (g *GetTransferSettingsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		TransferEnabled  types.Optional[bool]   `json:"transfer_enabled"`
		TransferInterval types.Optional[string] `json:"transfer_interval"`
		TransferDay      types.Optional[int]    `json:"transfer_day"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.TransferEnabled = temp.TransferEnabled
	g.TransferInterval = temp.TransferInterval
	g.TransferDay = temp.TransferDay
	return nil
}
