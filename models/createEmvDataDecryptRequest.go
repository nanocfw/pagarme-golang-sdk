package models

import (
	"encoding/json"
)

type CreateEmvDataDecryptRequest struct {
	Cipher string                            `json:"cipher"`
	Dukpt  *CreateEmvDataDukptDecryptRequest `json:"dukpt,omitempty"`
	Tags   []CreateEmvDataTlvDecryptRequest  `json:"tags"`
}

func (c *CreateEmvDataDecryptRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateEmvDataDecryptRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["cipher"] = c.Cipher
	if c.Dukpt != nil {
		structMap["dukpt"] = c.Dukpt
	}
	structMap["tags"] = c.Tags
	return structMap
}

func (c *CreateEmvDataDecryptRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Cipher string                            `json:"cipher"`
		Dukpt  *CreateEmvDataDukptDecryptRequest `json:"dukpt,omitempty"`
		Tags   []CreateEmvDataTlvDecryptRequest  `json:"tags"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Cipher = temp.Cipher
	c.Dukpt = temp.Dukpt
	c.Tags = temp.Tags
	return nil
}
