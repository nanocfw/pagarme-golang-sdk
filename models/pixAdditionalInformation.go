package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Pix Additional Information
type PixAdditionalInformation struct {
	Name  types.Optional[string] `json:"Name"`
	Value types.Optional[string] `json:"Value"`
}

func (p *PixAdditionalInformation) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

func (p *PixAdditionalInformation) toMap() map[string]any {
	structMap := make(map[string]any)
	if p.Name.IsValueSet() {
		structMap["Name"] = p.Name.Value()
	}
	if p.Value.IsValueSet() {
		structMap["Value"] = p.Value.Value()
	}
	return structMap
}

func (p *PixAdditionalInformation) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name  types.Optional[string] `json:"Name"`
		Value types.Optional[string] `json:"Value"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.Name = temp.Name
	p.Value = temp.Value
	return nil
}
