package models

import (
	"encoding/json"
	"log"
	"time"
)

type CreateCheckoutBoletoPaymentRequest struct {
	Bank         string    `json:"bank"`
	Instructions string    `json:"instructions"`
	DueAt        time.Time `json:"due_at"`
}

func (c *CreateCheckoutBoletoPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCheckoutBoletoPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["bank"] = c.Bank
	structMap["instructions"] = c.Instructions
	structMap["due_at"] = c.DueAt.Format(time.RFC3339)
	return structMap
}

func (c *CreateCheckoutBoletoPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Bank         string `json:"bank"`
		Instructions string `json:"instructions"`
		DueAt        string `json:"due_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Bank = temp.Bank
	c.Instructions = temp.Instructions
	DueAtVal, err := time.Parse(time.RFC3339, temp.DueAt)
	if err != nil {
		log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
	}
	c.DueAt = DueAtVal
	return nil
}
