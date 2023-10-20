package models

import (
	"encoding/json"
)

type GetCheckoutCardInstallmentOptionsResponse struct {
	Number *string `json:"number"`
	Total  *int    `json:"total"`
}

func (g *GetCheckoutCardInstallmentOptionsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetCheckoutCardInstallmentOptionsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["number"] = g.Number
	structMap["total"] = g.Total
	return structMap
}

func (g *GetCheckoutCardInstallmentOptionsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Number *string `json:"number"`
		Total  *int    `json:"total"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Number = temp.Number
	g.Total = temp.Total
	return nil
}
