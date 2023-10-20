package models

import (
	"encoding/json"
	"log"
	"time"
)

// Request for creating an anticipation
type CreateAnticipationRequest struct {
	Amount      int       `json:"amount"`
	Timeframe   string    `json:"timeframe"`
	PaymentDate time.Time `json:"payment_date"`
}

func (c *CreateAnticipationRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateAnticipationRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = c.Amount
	structMap["timeframe"] = c.Timeframe
	structMap["payment_date"] = c.PaymentDate.Format(time.RFC3339)
	return structMap
}

func (c *CreateAnticipationRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount      int    `json:"amount"`
		Timeframe   string `json:"timeframe"`
		PaymentDate string `json:"payment_date"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Amount = temp.Amount
	c.Timeframe = temp.Timeframe
	PaymentDateVal, err := time.Parse(time.RFC3339, temp.PaymentDate)
	if err != nil {
		log.Fatalf("Cannot Parse payment_date as % s format.", time.RFC3339)
	}
	c.PaymentDate = PaymentDateVal
	return nil
}
