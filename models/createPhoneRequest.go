package models

import (
	"encoding/json"
)

type CreatePhoneRequest struct {
	CountryCode *string `json:"country_code,omitempty"`
	Number      *string `json:"number,omitempty"`
	AreaCode    *string `json:"area_code,omitempty"`
}

func (c *CreatePhoneRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreatePhoneRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.CountryCode != nil {
		structMap["country_code"] = c.CountryCode
	}
	if c.Number != nil {
		structMap["number"] = c.Number
	}
	if c.AreaCode != nil {
		structMap["area_code"] = c.AreaCode
	}
	return structMap
}

func (c *CreatePhoneRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CountryCode *string `json:"country_code,omitempty"`
		Number      *string `json:"number,omitempty"`
		AreaCode    *string `json:"area_code,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.CountryCode = temp.CountryCode
	c.Number = temp.Number
	c.AreaCode = temp.AreaCode
	return nil
}
