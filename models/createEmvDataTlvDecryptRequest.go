package models

import (
	"encoding/json"
)

type CreateEmvDataTlvDecryptRequest struct {
	Tag    string `json:"tag"`
	Lenght string `json:"lenght"`
	Value  string `json:"value"`
}

func (c *CreateEmvDataTlvDecryptRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateEmvDataTlvDecryptRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["tag"] = c.Tag
	structMap["lenght"] = c.Lenght
	structMap["value"] = c.Value
	return structMap
}

func (c *CreateEmvDataTlvDecryptRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Tag    string `json:"tag"`
		Lenght string `json:"lenght"`
		Value  string `json:"value"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Tag = temp.Tag
	c.Lenght = temp.Lenght
	c.Value = temp.Value
	return nil
}
