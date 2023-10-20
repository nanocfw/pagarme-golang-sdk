package models

import (
	"encoding/json"
)

// SubMerchant
type CreateSubMerchantRequest struct {
	PaymentFacilitatorCode string               `json:"payment_facilitator_code"`
	Code                   string               `json:"code"`
	Name                   string               `json:"name"`
	MerchantCategoryCode   string               `json:"merchant_category_code"`
	Document               string               `json:"document"`
	Type                   string               `json:"type"`
	Phone                  CreatePhoneRequest   `json:"phone"`
	Address                CreateAddressRequest `json:"address"`
}

func (c *CreateSubMerchantRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateSubMerchantRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["payment_facilitator_code"] = c.PaymentFacilitatorCode
	structMap["code"] = c.Code
	structMap["name"] = c.Name
	structMap["merchant_category_code"] = c.MerchantCategoryCode
	structMap["document"] = c.Document
	structMap["type"] = c.Type
	structMap["phone"] = c.Phone
	structMap["address"] = c.Address
	return structMap
}

func (c *CreateSubMerchantRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PaymentFacilitatorCode string               `json:"payment_facilitator_code"`
		Code                   string               `json:"code"`
		Name                   string               `json:"name"`
		MerchantCategoryCode   string               `json:"merchant_category_code"`
		Document               string               `json:"document"`
		Type                   string               `json:"type"`
		Phone                  CreatePhoneRequest   `json:"phone"`
		Address                CreateAddressRequest `json:"address"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.PaymentFacilitatorCode = temp.PaymentFacilitatorCode
	c.Code = temp.Code
	c.Name = temp.Name
	c.MerchantCategoryCode = temp.MerchantCategoryCode
	c.Document = temp.Document
	c.Type = temp.Type
	c.Phone = temp.Phone
	c.Address = temp.Address
	return nil
}
