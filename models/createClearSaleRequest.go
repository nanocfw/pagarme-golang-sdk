package models

import (
	"encoding/json"
)

type CreateClearSaleRequest struct {
	CustomSla int `json:"custom_sla"`
}

func (c *CreateClearSaleRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateClearSaleRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["custom_sla"] = c.CustomSla
	return structMap
}

func (c *CreateClearSaleRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CustomSla int `json:"custom_sla"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.CustomSla = temp.CustomSla
	return nil
}
