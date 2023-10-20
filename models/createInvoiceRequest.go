package models

import (
	"encoding/json"
)

// Request for creating a new Invoice
type CreateInvoiceRequest struct {
	Metadata map[string]string `json:"metadata"`
}

func (c *CreateInvoiceRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateInvoiceRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["metadata"] = c.Metadata
	return structMap
}

func (c *CreateInvoiceRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Metadata map[string]string `json:"metadata"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Metadata = temp.Metadata
	return nil
}
