package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

type GetPhoneResponse struct {
	CountryCode types.Optional[string] `json:"country_code"`
	Number      types.Optional[string] `json:"number"`
	AreaCode    types.Optional[string] `json:"area_code"`
}

func (g *GetPhoneResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetPhoneResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.CountryCode.IsValueSet() {
		structMap["country_code"] = g.CountryCode.Value()
	}
	if g.Number.IsValueSet() {
		structMap["number"] = g.Number.Value()
	}
	if g.AreaCode.IsValueSet() {
		structMap["area_code"] = g.AreaCode.Value()
	}
	return structMap
}

func (g *GetPhoneResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CountryCode types.Optional[string] `json:"country_code"`
		Number      types.Optional[string] `json:"number"`
		AreaCode    types.Optional[string] `json:"area_code"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.CountryCode = temp.CountryCode
	g.Number = temp.Number
	g.AreaCode = temp.AreaCode
	return nil
}
