package models

import (
	"encoding/json"
)

type CreateCashPaymentRequest struct {
	Description string `json:"description"`
	Confirm     bool   `json:"confirm"`
}

func (c *CreateCashPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCashPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["description"] = c.Description
	structMap["confirm"] = c.Confirm
	return structMap
}

func (c *CreateCashPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Description string `json:"description"`
		Confirm     bool   `json:"confirm"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Description = temp.Description
	c.Confirm = temp.Confirm
	return nil
}
