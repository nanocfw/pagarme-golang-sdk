package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for geetting an order device
type GetDeviceResponse struct {
	Platform types.Optional[string] `json:"platform"`
}

func (g *GetDeviceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetDeviceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Platform.IsValueSet() {
		structMap["platform"] = g.Platform.Value()
	}
	return structMap
}

func (g *GetDeviceResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Platform types.Optional[string] `json:"platform"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Platform = temp.Platform
	return nil
}
