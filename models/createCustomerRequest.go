package models

import (
	"encoding/json"
)

// Request for creating a new customer
type CreateCustomerRequest struct {
	Name         string               `json:"name"`
	Email        string               `json:"email"`
	Document     string               `json:"document"`
	Type         string               `json:"type"`
	Address      CreateAddressRequest `json:"address"`
	Metadata     map[string]string    `json:"metadata,omitempty"`
	Phones       CreatePhonesRequest  `json:"phones"`
	Code         string               `json:"code,omitempty"`
	Gender       *string              `json:"gender,omitempty"`
	DocumentType *string              `json:"document_type,omitempty"`
}

func (c *CreateCustomerRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateCustomerRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["name"] = c.Name
	structMap["email"] = c.Email
	structMap["document"] = c.Document
	structMap["type"] = c.Type
	structMap["address"] = c.Address
	structMap["metadata"] = c.Metadata
	structMap["phones"] = c.Phones
	structMap["code"] = c.Code
	if c.Gender != nil {
		structMap["gender"] = c.Gender
	}
	if c.DocumentType != nil {
		structMap["document_type"] = c.DocumentType
	}
	return structMap
}

func (c *CreateCustomerRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name         string               `json:"name"`
		Email        string               `json:"email"`
		Document     string               `json:"document"`
		Type         string               `json:"type"`
		Address      CreateAddressRequest `json:"address"`
		Metadata     map[string]string    `json:"metadata"`
		Phones       CreatePhonesRequest  `json:"phones"`
		Code         string               `json:"code"`
		Gender       *string              `json:"gender,omitempty"`
		DocumentType *string              `json:"document_type,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Name = temp.Name
	c.Email = temp.Email
	c.Document = temp.Document
	c.Type = temp.Type
	c.Address = temp.Address
	c.Metadata = temp.Metadata
	c.Phones = temp.Phones
	c.Code = temp.Code
	c.Gender = temp.Gender
	c.DocumentType = temp.DocumentType
	return nil
}
