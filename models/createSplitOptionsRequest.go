package models

import (
	"encoding/json"
)

// The Split Options Request
type CreateSplitOptionsRequest struct {
	Liable              *bool `json:"liable,omitempty"`
	ChargeProcessingFee *bool `json:"charge_processing_fee,omitempty"`
	ChargeRemainderFee  *bool `json:"charge_remainder_fee,omitempty"`
}

func (c *CreateSplitOptionsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateSplitOptionsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Liable != nil {
		structMap["liable"] = c.Liable
	}
	if c.ChargeProcessingFee != nil {
		structMap["charge_processing_fee"] = c.ChargeProcessingFee
	}
	if c.ChargeRemainderFee != nil {
		structMap["charge_remainder_fee"] = c.ChargeRemainderFee
	}
	return structMap
}

func (c *CreateSplitOptionsRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Liable              *bool `json:"liable,omitempty"`
		ChargeProcessingFee *bool `json:"charge_processing_fee,omitempty"`
		ChargeRemainderFee  *bool `json:"charge_remainder_fee,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Liable = temp.Liable
	c.ChargeProcessingFee = temp.ChargeProcessingFee
	c.ChargeRemainderFee = temp.ChargeRemainderFee
	return nil
}
