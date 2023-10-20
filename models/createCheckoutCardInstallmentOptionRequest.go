package models

import (
	"encoding/json"
)

// Options for card installment
type CreateCheckoutCardInstallmentOptionRequest struct {
	Number int `json:"number"`
	Total  int `json:"total"`
}

func (c *CreateCheckoutCardInstallmentOptionRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCheckoutCardInstallmentOptionRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["number"] = c.Number
	structMap["total"] = c.Total
	return structMap
}

func (c *CreateCheckoutCardInstallmentOptionRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Number int `json:"number"`
		Total  int `json:"total"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Number = temp.Number
	c.Total = temp.Total
	return nil
}
