package models

import (
	"encoding/json"
)

// Card token data
type CreateCardTokenRequest struct {
	Number     string `json:"number"`
	HolderName string `json:"holder_name"`
	ExpMonth   int    `json:"exp_month"`
	ExpYear    int    `json:"exp_year"`
	Cvv        string `json:"cvv"`
	Brand      string `json:"brand"`
	Label      string `json:"label"`
}

func (c *CreateCardTokenRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCardTokenRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["number"] = c.Number
	structMap["holder_name"] = c.HolderName
	structMap["exp_month"] = c.ExpMonth
	structMap["exp_year"] = c.ExpYear
	structMap["cvv"] = c.Cvv
	structMap["brand"] = c.Brand
	structMap["label"] = c.Label
	return structMap
}

func (c *CreateCardTokenRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Number     string `json:"number"`
		HolderName string `json:"holder_name"`
		ExpMonth   int    `json:"exp_month"`
		ExpYear    int    `json:"exp_year"`
		Cvv        string `json:"cvv"`
		Brand      string `json:"brand"`
		Label      string `json:"label"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Number = temp.Number
	c.HolderName = temp.HolderName
	c.ExpMonth = temp.ExpMonth
	c.ExpYear = temp.ExpYear
	c.Cvv = temp.Cvv
	c.Brand = temp.Brand
	c.Label = temp.Label
	return nil
}
