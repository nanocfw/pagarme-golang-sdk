package models

import (
	"encoding/json"
)

type CreateEmvDataDukptDecryptRequest struct {
	Ksn string `json:"ksn"`
}

func (c *CreateEmvDataDukptDecryptRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateEmvDataDukptDecryptRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["ksn"] = c.Ksn
	return structMap
}

func (c *CreateEmvDataDukptDecryptRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Ksn string `json:"ksn"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Ksn = temp.Ksn
	return nil
}
