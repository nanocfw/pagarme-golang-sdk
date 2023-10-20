package models

import (
	"encoding/json"
)

// Request for capturing a charge
type CreateCaptureChargeRequest struct {
	Code               string                `json:"code"`
	Amount             *int                  `json:"amount,omitempty"`
	Split              *[]CreateSplitRequest `json:"split,omitempty"`
	OperationReference string                `json:"operation_reference"`
}

func (c *CreateCaptureChargeRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCaptureChargeRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["code"] = c.Code
	if c.Amount != nil {
		structMap["amount"] = c.Amount
	}
	if c.Split != nil {
		structMap["split"] = c.Split
	}
	structMap["operation_reference"] = c.OperationReference
	return structMap
}

func (c *CreateCaptureChargeRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Code               string                `json:"code"`
		Amount             *int                  `json:"amount,omitempty"`
		Split              *[]CreateSplitRequest `json:"split,omitempty"`
		OperationReference string                `json:"operation_reference"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Code = temp.Code
	c.Amount = temp.Amount
	c.Split = temp.Split
	c.OperationReference = temp.OperationReference
	return nil
}
