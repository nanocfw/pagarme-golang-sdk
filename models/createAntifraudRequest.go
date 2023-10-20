package models

import (
	"encoding/json"
)

type CreateAntifraudRequest struct {
	Type      string                 `json:"type"`
	Clearsale CreateClearSaleRequest `json:"clearsale"`
}

func (c *CreateAntifraudRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateAntifraudRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["type"] = c.Type
	structMap["clearsale"] = c.Clearsale
	return structMap
}

func (c *CreateAntifraudRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Type      string                 `json:"type"`
		Clearsale CreateClearSaleRequest `json:"clearsale"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Type = temp.Type
	c.Clearsale = temp.Clearsale
	return nil
}
