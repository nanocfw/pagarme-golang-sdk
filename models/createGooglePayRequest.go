package models

import (
	"encoding/json"
)

// The GooglePay Token Payment Request
type CreateGooglePayRequest struct {
	Version            string                       `json:"version"`
	Data               string                       `json:"data"`
	Header             CreateGooglePayHeaderRequest `json:"header"`
	Signature          string                       `json:"signature"`
	MerchantIdentifier string                       `json:"merchant_identifier"`
}

func (c *CreateGooglePayRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateGooglePayRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["version"] = c.Version
	structMap["data"] = c.Data
	structMap["header"] = c.Header
	structMap["signature"] = c.Signature
	structMap["merchant_identifier"] = c.MerchantIdentifier
	return structMap
}

func (c *CreateGooglePayRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Version            string                       `json:"version"`
		Data               string                       `json:"data"`
		Header             CreateGooglePayHeaderRequest `json:"header"`
		Signature          string                       `json:"signature"`
		MerchantIdentifier string                       `json:"merchant_identifier"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Version = temp.Version
	c.Data = temp.Data
	c.Header = temp.Header
	c.Signature = temp.Signature
	c.MerchantIdentifier = temp.MerchantIdentifier
	return nil
}
