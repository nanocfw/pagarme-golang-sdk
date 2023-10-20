package models

import (
	"encoding/json"
)

// Creates a 3D-S authentication payment
type CreateThreeDSecureRequest struct {
	Mpi             string  `json:"mpi"`
	Cavv            *string `json:"cavv,omitempty"`
	Eci             *string `json:"eci,omitempty"`
	TransactionId   *string `json:"transaction_id,omitempty"`
	SuccessUrl      *string `json:"success_url,omitempty"`
	DsTransactionId *string `json:"ds_transaction_id,omitempty"`
	Version         *string `json:"version,omitempty"`
}

func (c *CreateThreeDSecureRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateThreeDSecureRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["mpi"] = c.Mpi
	if c.Cavv != nil {
		structMap["cavv"] = c.Cavv
	}
	if c.Eci != nil {
		structMap["eci"] = c.Eci
	}
	if c.TransactionId != nil {
		structMap["transaction_id"] = c.TransactionId
	}
	if c.SuccessUrl != nil {
		structMap["success_url"] = c.SuccessUrl
	}
	if c.DsTransactionId != nil {
		structMap["ds_transaction_id"] = c.DsTransactionId
	}
	if c.Version != nil {
		structMap["version"] = c.Version
	}
	return structMap
}

func (c *CreateThreeDSecureRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Mpi             string  `json:"mpi"`
		Cavv            *string `json:"cavv,omitempty"`
		Eci             *string `json:"eci,omitempty"`
		TransactionId   *string `json:"transaction_id,omitempty"`
		SuccessUrl      *string `json:"success_url,omitempty"`
		DsTransactionId *string `json:"ds_transaction_id,omitempty"`
		Version         *string `json:"version,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Mpi = temp.Mpi
	c.Cavv = temp.Cavv
	c.Eci = temp.Eci
	c.TransactionId = temp.TransactionId
	c.SuccessUrl = temp.SuccessUrl
	c.DsTransactionId = temp.DsTransactionId
	c.Version = temp.Version
	return nil
}
