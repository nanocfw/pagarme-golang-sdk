package models

import (
	"encoding/json"
)

type CreatePhonesRequest struct {
	HomePhone   *CreatePhoneRequest `json:"home_phone,omitempty"`
	MobilePhone *CreatePhoneRequest `json:"mobile_phone,omitempty"`
}

func (c *CreatePhonesRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreatePhonesRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.HomePhone != nil {
		structMap["home_phone"] = c.HomePhone
	}
	if c.MobilePhone != nil {
		structMap["mobile_phone"] = c.MobilePhone
	}
	return structMap
}

func (c *CreatePhonesRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		HomePhone   *CreatePhoneRequest `json:"home_phone,omitempty"`
		MobilePhone *CreatePhoneRequest `json:"mobile_phone,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.HomePhone = temp.HomePhone
	c.MobilePhone = temp.MobilePhone
	return nil
}
