package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Response object for geetting an order location request
type GetLocationResponse struct {
	Latitude  types.Optional[string] `json:"latitude"`
	Longitude types.Optional[string] `json:"longitude"`
}

func (g *GetLocationResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetLocationResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Latitude.IsValueSet() {
		structMap["latitude"] = g.Latitude.Value()
	}
	if g.Longitude.IsValueSet() {
		structMap["longitude"] = g.Longitude.Value()
	}
	return structMap
}

func (g *GetLocationResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Latitude  types.Optional[string] `json:"latitude"`
		Longitude types.Optional[string] `json:"longitude"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Latitude = temp.Latitude
	g.Longitude = temp.Longitude
	return nil
}
