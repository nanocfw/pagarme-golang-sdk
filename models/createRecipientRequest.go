package models

import (
	"encoding/json"
)

// Request for creating a recipient
type CreateRecipientRequest struct {
	Name               string                         `json:"name"`
	Email              string                         `json:"email"`
	Description        string                         `json:"description"`
	Document           string                         `json:"document"`
	Type               string                         `json:"type"`
	DefaultBankAccount CreateBankAccountRequest       `json:"default_bank_account"`
	Metadata           map[string]string              `json:"metadata"`
	TransferSettings   *CreateTransferSettingsRequest `json:"transfer_settings,omitempty"`
	Code               string                         `json:"code"`
	PaymentMode        string                         `json:"payment_mode"`
}

func (c *CreateRecipientRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateRecipientRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["name"] = c.Name
	structMap["email"] = c.Email
	structMap["description"] = c.Description
	structMap["document"] = c.Document
	structMap["type"] = c.Type
	structMap["default_bank_account"] = c.DefaultBankAccount
	structMap["metadata"] = c.Metadata
	if c.TransferSettings != nil {
		structMap["transfer_settings"] = c.TransferSettings
	}
	structMap["code"] = c.Code
	structMap["payment_mode"] = c.PaymentMode
	return structMap
}

func (c *CreateRecipientRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name               string                         `json:"name"`
		Email              string                         `json:"email"`
		Description        string                         `json:"description"`
		Document           string                         `json:"document"`
		Type               string                         `json:"type"`
		DefaultBankAccount CreateBankAccountRequest       `json:"default_bank_account"`
		Metadata           map[string]string              `json:"metadata"`
		TransferSettings   *CreateTransferSettingsRequest `json:"transfer_settings,omitempty"`
		Code               string                         `json:"code"`
		PaymentMode        string                         `json:"payment_mode"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Name = temp.Name
	c.Email = temp.Email
	c.Description = temp.Description
	c.Document = temp.Document
	c.Type = temp.Type
	c.DefaultBankAccount = temp.DefaultBankAccount
	c.Metadata = temp.Metadata
	c.TransferSettings = temp.TransferSettings
	c.Code = temp.Code
	c.PaymentMode = temp.PaymentMode
	return nil
}
