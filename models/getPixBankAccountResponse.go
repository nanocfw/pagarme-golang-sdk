package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

// Payer's bank details.
type GetPixBankAccountResponse struct {
	BankName      types.Optional[string] `json:"bank_name"`
	Ispb          types.Optional[string] `json:"ispb"`
	BranchCode    types.Optional[string] `json:"branch_code"`
	AccountNumber types.Optional[string] `json:"account_number"`
}

func (g *GetPixBankAccountResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetPixBankAccountResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.BankName.IsValueSet() {
		structMap["bank_name"] = g.BankName.Value()
	}
	if g.Ispb.IsValueSet() {
		structMap["ispb"] = g.Ispb.Value()
	}
	if g.BranchCode.IsValueSet() {
		structMap["branch_code"] = g.BranchCode.Value()
	}
	if g.AccountNumber.IsValueSet() {
		structMap["account_number"] = g.AccountNumber.Value()
	}
	return structMap
}

func (g *GetPixBankAccountResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		BankName      types.Optional[string] `json:"bank_name"`
		Ispb          types.Optional[string] `json:"ispb"`
		BranchCode    types.Optional[string] `json:"branch_code"`
		AccountNumber types.Optional[string] `json:"account_number"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.BankName = temp.BankName
	g.Ispb = temp.Ispb
	g.BranchCode = temp.BranchCode
	g.AccountNumber = temp.AccountNumber
	return nil
}
