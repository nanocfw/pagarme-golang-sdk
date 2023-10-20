package models

import (
	"encoding/json"
)

// Card data
type CreateCardRequest struct {
	Number           *string                   `json:"number,omitempty"`
	HolderName       *string                   `json:"holder_name,omitempty"`
	ExpMonth         *int                      `json:"exp_month,omitempty"`
	ExpYear          *int                      `json:"exp_year,omitempty"`
	Cvv              *string                   `json:"cvv,omitempty"`
	BillingAddress   *CreateAddressRequest     `json:"billing_address,omitempty"`
	Brand            *string                   `json:"brand,omitempty"`
	BillingAddressId *string                   `json:"billing_address_id,omitempty"`
	Metadata         map[string]*string        `json:"metadata,omitempty"`
	Type             *string                   `json:"type,omitempty"`
	Options          *CreateCardOptionsRequest `json:"options,omitempty"`
	HolderDocument   *string                   `json:"holder_document,omitempty"`
	PrivateLabel     *bool                     `json:"private_label,omitempty"`
	Label            *string                   `json:"label,omitempty"`
	Id               *string                   `json:"id,omitempty"`
	Token            *string                   `json:"token,omitempty"`
}

func (c *CreateCardRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCardRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Number != nil {
		structMap["number"] = c.Number
	}
	if c.HolderName != nil {
		structMap["holder_name"] = c.HolderName
	}
	if c.ExpMonth != nil {
		structMap["exp_month"] = c.ExpMonth
	}
	if c.ExpYear != nil {
		structMap["exp_year"] = c.ExpYear
	}
	if c.Cvv != nil {
		structMap["cvv"] = c.Cvv
	}
	if c.BillingAddress != nil {
		structMap["billing_address"] = c.BillingAddress
	}
	if c.Brand != nil {
		structMap["brand"] = c.Brand
	}
	if c.BillingAddressId != nil {
		structMap["billing_address_id"] = c.BillingAddressId
	}
	if c.Metadata != nil {
		structMap["metadata"] = c.Metadata
	}
	if c.Type != nil {
		structMap["type"] = c.Type
	}
	if c.Options != nil {
		structMap["options"] = c.Options
	}
	if c.HolderDocument != nil {
		structMap["holder_document"] = c.HolderDocument
	}
	if c.PrivateLabel != nil {
		structMap["private_label"] = c.PrivateLabel
	}
	if c.Label != nil {
		structMap["label"] = c.Label
	}
	if c.Id != nil {
		structMap["id"] = c.Id
	}
	if c.Token != nil {
		structMap["token"] = c.Token
	}
	return structMap
}

func (c *CreateCardRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Number           *string                   `json:"number,omitempty"`
		HolderName       *string                   `json:"holder_name,omitempty"`
		ExpMonth         *int                      `json:"exp_month,omitempty"`
		ExpYear          *int                      `json:"exp_year,omitempty"`
		Cvv              *string                   `json:"cvv,omitempty"`
		BillingAddress   *CreateAddressRequest     `json:"billing_address,omitempty"`
		Brand            *string                   `json:"brand,omitempty"`
		BillingAddressId *string                   `json:"billing_address_id,omitempty"`
		Metadata         map[string]*string        `json:"metadata,omitempty"`
		Type             *string                   `json:"type,omitempty"`
		Options          *CreateCardOptionsRequest `json:"options,omitempty"`
		HolderDocument   *string                   `json:"holder_document,omitempty"`
		PrivateLabel     *bool                     `json:"private_label,omitempty"`
		Label            *string                   `json:"label,omitempty"`
		Id               *string                   `json:"id,omitempty"`
		Token            *string                   `json:"token,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Number = temp.Number
	c.HolderName = temp.HolderName
	c.ExpMonth = temp.ExpMonth
	c.ExpYear = temp.ExpYear
	c.Cvv = temp.Cvv
	c.BillingAddress = temp.BillingAddress
	c.Brand = temp.Brand
	c.BillingAddressId = temp.BillingAddressId
	c.Metadata = temp.Metadata
	c.Type = temp.Type
	c.Options = temp.Options
	c.HolderDocument = temp.HolderDocument
	c.PrivateLabel = temp.PrivateLabel
	c.Label = temp.Label
	c.Id = temp.Id
	c.Token = temp.Token
	return nil
}
