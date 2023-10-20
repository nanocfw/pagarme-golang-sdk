package models

import (
	"encoding/json"
	"log"
	"time"
)

// Request for creating a usage
type CreateUsageRequest struct {
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
	UsedAt      time.Time `json:"used_at"`
	Code        *string   `json:"code,omitempty"`
	Group       *string   `json:"group,omitempty"`
	Amount      *int      `json:"amount,omitempty"`
}

func (c *CreateUsageRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateUsageRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["quantity"] = c.Quantity
	structMap["description"] = c.Description
	structMap["used_at"] = c.UsedAt.Format(time.RFC3339)
	if c.Code != nil {
		structMap["code"] = c.Code
	}
	if c.Group != nil {
		structMap["group"] = c.Group
	}
	if c.Amount != nil {
		structMap["amount"] = c.Amount
	}
	return structMap
}

func (c *CreateUsageRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Quantity    int     `json:"quantity"`
		Description string  `json:"description"`
		UsedAt      string  `json:"used_at"`
		Code        *string `json:"code,omitempty"`
		Group       *string `json:"group,omitempty"`
		Amount      *int    `json:"amount,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Quantity = temp.Quantity
	c.Description = temp.Description
	UsedAtVal, err := time.Parse(time.RFC3339, temp.UsedAt)
	if err != nil {
		log.Fatalf("Cannot Parse used_at as % s format.", time.RFC3339)
	}
	c.UsedAt = UsedAtVal
	c.Code = temp.Code
	c.Group = temp.Group
	c.Amount = temp.Amount
	return nil
}
