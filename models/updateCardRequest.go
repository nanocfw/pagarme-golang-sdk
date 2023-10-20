package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Request for updating a card
type UpdateCardRequest struct {
	HolderName       string                 `json:"holder_name"`
	ExpMonth         int                    `json:"exp_month"`
	ExpYear          int                    `json:"exp_year"`
	BillingAddressId types.Optional[string] `json:"billing_address_id"`
	BillingAddress   CreateAddressRequest   `json:"billing_address"`
	Metadata         map[string]string      `json:"metadata"`
	Label            string                 `json:"label"`
}

func (u *UpdateCardRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

func (u *UpdateCardRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["holder_name"] = u.HolderName
	structMap["exp_month"] = u.ExpMonth
	structMap["exp_year"] = u.ExpYear
	if u.BillingAddressId.IsValueSet() {
		structMap["billing_address_id"] = u.BillingAddressId.Value()
	}
	structMap["billing_address"] = u.BillingAddress
	structMap["metadata"] = u.Metadata
	structMap["label"] = u.Label
	return structMap
}

func (u *UpdateCardRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		HolderName       string                 `json:"holder_name"`
		ExpMonth         int                    `json:"exp_month"`
		ExpYear          int                    `json:"exp_year"`
		BillingAddressId types.Optional[string] `json:"billing_address_id"`
		BillingAddress   CreateAddressRequest   `json:"billing_address"`
		Metadata         map[string]string      `json:"metadata"`
		Label            string                 `json:"label"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.HolderName = temp.HolderName
	u.ExpMonth = temp.ExpMonth
	u.ExpYear = temp.ExpYear
	u.BillingAddressId = temp.BillingAddressId
	u.BillingAddress = temp.BillingAddress
	u.Metadata = temp.Metadata
	u.Label = temp.Label
	return nil
}
