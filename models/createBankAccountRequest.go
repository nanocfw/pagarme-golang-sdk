package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Request for creating a bank account
type CreateBankAccountRequest struct {
	HolderName        string                 `json:"holder_name"`
	HolderType        string                 `json:"holder_type"`
	HolderDocument    string                 `json:"holder_document"`
	Bank              string                 `json:"bank"`
	BranchNumber      string                 `json:"branch_number"`
	BranchCheckDigit  types.Optional[string] `json:"branch_check_digit"`
	AccountNumber     string                 `json:"account_number"`
	AccountCheckDigit string                 `json:"account_check_digit"`
	Type              string                 `json:"type"`
	Metadata          map[string]string      `json:"metadata"`
	PixKey            types.Optional[string] `json:"pix_key"`
}

func (c *CreateBankAccountRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

func (c *CreateBankAccountRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["holder_name"] = c.HolderName
	structMap["holder_type"] = c.HolderType
	structMap["holder_document"] = c.HolderDocument
	structMap["bank"] = c.Bank
	structMap["branch_number"] = c.BranchNumber
	if c.BranchCheckDigit.IsValueSet() {
		structMap["branch_check_digit"] = c.BranchCheckDigit.Value()
	}
	structMap["account_number"] = c.AccountNumber
	structMap["account_check_digit"] = c.AccountCheckDigit
	structMap["type"] = c.Type
	structMap["metadata"] = c.Metadata
	if c.PixKey.IsValueSet() {
		structMap["pix_key"] = c.PixKey.Value()
	}
	return structMap
}

func (c *CreateBankAccountRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		HolderName        string                 `json:"holder_name"`
		HolderType        string                 `json:"holder_type"`
		HolderDocument    string                 `json:"holder_document"`
		Bank              string                 `json:"bank"`
		BranchNumber      string                 `json:"branch_number"`
		BranchCheckDigit  types.Optional[string] `json:"branch_check_digit"`
		AccountNumber     string                 `json:"account_number"`
		AccountCheckDigit string                 `json:"account_check_digit"`
		Type              string                 `json:"type"`
		Metadata          map[string]string      `json:"metadata"`
		PixKey            types.Optional[string] `json:"pix_key"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.HolderName = temp.HolderName
	c.HolderType = temp.HolderType
	c.HolderDocument = temp.HolderDocument
	c.Bank = temp.Bank
	c.BranchNumber = temp.BranchNumber
	c.BranchCheckDigit = temp.BranchCheckDigit
	c.AccountNumber = temp.AccountNumber
	c.AccountCheckDigit = temp.AccountCheckDigit
	c.Type = temp.Type
	c.Metadata = temp.Metadata
	c.PixKey = temp.PixKey
	return nil
}
