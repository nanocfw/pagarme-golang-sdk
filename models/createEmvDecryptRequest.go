package models

import (
	"encoding/json"
)

type CreateEmvDecryptRequest struct {
	IccData            string                                  `json:"icc_data"`
	CardSequenceNumber string                                  `json:"card_sequence_number"`
	Data               CreateEmvDataDecryptRequest             `json:"data"`
	Poi                *CreateCardPaymentContactlessPOIRequest `json:"poi,omitempty"`
}

func (c *CreateEmvDecryptRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateEmvDecryptRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["icc_data"] = c.IccData
	structMap["card_sequence_number"] = c.CardSequenceNumber
	structMap["data"] = c.Data
	if c.Poi != nil {
		structMap["poi"] = c.Poi
	}
	return structMap
}

func (c *CreateEmvDecryptRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		IccData            string                                  `json:"icc_data"`
		CardSequenceNumber string                                  `json:"card_sequence_number"`
		Data               CreateEmvDataDecryptRequest             `json:"data"`
		Poi                *CreateCardPaymentContactlessPOIRequest `json:"poi,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.IccData = temp.IccData
	c.CardSequenceNumber = temp.CardSequenceNumber
	c.Data = temp.Data
	c.Poi = temp.Poi
	return nil
}
