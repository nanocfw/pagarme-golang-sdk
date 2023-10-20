package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

type GetPhonesResponse struct {
	HomePhone   types.Optional[GetPhoneResponse] `json:"home_phone"`
	MobilePhone types.Optional[GetPhoneResponse] `json:"mobile_phone"`
}

func (g *GetPhonesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetPhonesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.HomePhone.IsValueSet() {
		structMap["home_phone"] = g.HomePhone.Value()
	}
	if g.MobilePhone.IsValueSet() {
		structMap["mobile_phone"] = g.MobilePhone.Value()
	}
	return structMap
}

func (g *GetPhonesResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		HomePhone   types.Optional[GetPhoneResponse] `json:"home_phone"`
		MobilePhone types.Optional[GetPhoneResponse] `json:"mobile_phone"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.HomePhone = temp.HomePhone
	g.MobilePhone = temp.MobilePhone
	return nil
}
